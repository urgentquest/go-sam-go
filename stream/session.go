package stream

import (
	"context"
	"net"
	"time"

	"github.com/go-i2p/go-sam-go/common"
	"github.com/go-i2p/i2pkeys"
)

// Read reads data from the stream.
func (s *StreamSession) Read(buf []byte) (int, error) {
	return s.Conn.Read(buf)
}

// Write sends data over the stream.
func (s *StreamSession) Write(data []byte) (int, error) {
	return s.Conn.Write(data)
}

func (s *StreamSession) SetDeadline(t time.Time) error {
	log.WithField("deadline", t).Debug("Setting deadline for StreamSession")
	return s.Conn.SetDeadline(t)
}

func (s *StreamSession) SetReadDeadline(t time.Time) error {
	log.WithField("readDeadline", t).Debug("Setting read deadline for StreamSession")
	return s.Conn.SetReadDeadline(t)
}

func (s *StreamSession) SetWriteDeadline(t time.Time) error {
	log.WithField("writeDeadline", t).Debug("Setting write deadline for StreamSession")
	return s.Conn.SetWriteDeadline(t)
}

func (s *StreamSession) From() string {
	return s.Fromport
}

func (s *StreamSession) To() string {
	return s.Toport
}

func (s *StreamSession) SignatureType() string {
	return s.SignatureType()
}

func (s *StreamSession) Close() error {
	log.WithField("id", s.ID()).Debug("Closing StreamSession")
	return s.Conn.Close()
}

// Returns the I2P destination (the address) of the stream session
func (s *StreamSession) Addr() i2pkeys.I2PAddr {
	return s.Addr()
}

func (s *StreamSession) LocalAddr() net.Addr {
	return s.Addr()
}

// Returns the keys associated with the stream session
func (s *StreamSession) Keys() i2pkeys.I2PKeys {
	return *s.DestinationKeys
}

// lookup name, convenience function
func (s *StreamSession) Lookup(name string) (i2pkeys.I2PAddr, error) {
	log.WithField("name", name).Debug("Looking up address")
	sam, err := common.NewSAM(s.Sam())
	if err == nil {
		addr, err := sam.Lookup(name)
		defer sam.Close()
		if err != nil {
			log.WithError(err).Error("Lookup failed")
		} else {
			log.WithField("addr", addr).Debug("Lookup successful")
		}
		return addr, err
	}
	log.WithError(err).Error("Failed to create SAM instance for lookup")
	return i2pkeys.I2PAddr(""), err
}

/*
func (s *StreamSession) Cancel() chan *StreamSession {
	ch := make(chan *StreamSession)
	ch <- s
	return ch
}*/

// deadline returns the earliest of:
//   - now+Timeout
//   - d.Deadline
//   - the context's deadline
//
// Or zero, if none of Timeout, Deadline, or context's deadline is set.
func (s *StreamSession) deadline(ctx context.Context, now time.Time) (earliest time.Time) {
	if s.Timeout != 0 { // including negative, for historical reasons
		earliest = now.Add(s.Timeout)
	}
	if d, ok := ctx.Deadline(); ok {
		earliest = minNonzeroTime(earliest, d)
	}
	return minNonzeroTime(earliest, s.Deadline)
}
