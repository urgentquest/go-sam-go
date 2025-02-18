package datagram

import (
	"bytes"
	"errors"
	"net"
	"time"

	"github.com/go-i2p/go-sam-go/common"
	"github.com/go-i2p/i2pkeys"
	"github.com/sirupsen/logrus"
)

func (s *DatagramSession) B32() string {
	b32 := s.DestinationKeys.Addr().Base32()
	log.WithField("b32", b32).Debug("Generated B32 address")
	return b32
}

func (s *DatagramSession) Dial(net string, addr string) (*DatagramSession, error) {
	log.WithFields(logrus.Fields{
		"net":  net,
		"addr": addr,
	}).Debug("Dialing address")
	netaddr, err := s.Lookup(addr)
	if err != nil {
		log.WithError(err).Error("Lookup failed")
		return nil, err
	}
	return s.DialI2PRemote(net, netaddr)
}

func (s *DatagramSession) DialRemote(net, addr string) (net.PacketConn, error) {
	log.WithFields(logrus.Fields{
		"net":  net,
		"addr": addr,
	}).Debug("Dialing remote address")
	netaddr, err := s.Lookup(addr)
	if err != nil {
		log.WithError(err).Error("Lookup failed")
		return nil, err
	}
	return s.DialI2PRemote(net, netaddr)
}

func (s *DatagramSession) DialI2PRemote(net string, addr net.Addr) (*DatagramSession, error) {
	log.WithFields(logrus.Fields{
		"net":  net,
		"addr": addr,
	}).Debug("Dialing I2P remote address")
	switch addr.(type) {
	case *i2pkeys.I2PAddr:
		s.RemoteI2PAddr = addr.(*i2pkeys.I2PAddr)
	case i2pkeys.I2PAddr:
		i2paddr := addr.(i2pkeys.I2PAddr)
		s.RemoteI2PAddr = &i2paddr
	}
	return s, nil
}

func (s *DatagramSession) RemoteAddr() net.Addr {
	log.WithField("remoteAddr", s.RemoteI2PAddr).Debug("Getting remote address")
	return s.RemoteI2PAddr
}

// Reads one datagram sent to the destination of the DatagramSession. Returns
// the number of bytes read, from what address it was sent, or an error.
// implements net.PacketConn
func (s *DatagramSession) ReadFrom(b []byte) (n int, addr net.Addr, err error) {
	log.Debug("Reading datagram")
	// extra bytes to read the remote address of incomming datagram
	buf := make([]byte, len(b)+4096)

	for {
		// very basic protection: only accept incomming UDP messages from the IP of the SAM bridge
		var saddr *net.UDPAddr
		n, saddr, err = s.UDPConn.ReadFromUDP(buf)
		if err != nil {
			log.WithError(err).Error("Failed to read from UDP")
			return 0, i2pkeys.I2PAddr(""), err
		}
		if bytes.Equal(saddr.IP, s.SAMUDPAddress.IP) {
			continue
		}
		break
	}
	i := bytes.IndexByte(buf, byte('\n'))
	if i > 4096 || i > n {
		log.Error("Could not parse incoming message remote address")
		return 0, i2pkeys.I2PAddr(""), errors.New("Could not parse incomming message remote address.")
	}
	raddr, err := i2pkeys.NewI2PAddrFromString(string(buf[:i]))
	if err != nil {
		log.WithError(err).Error("Could not parse incoming message remote address")
		return 0, i2pkeys.I2PAddr(""), errors.New("Could not parse incomming message remote address: " + err.Error())
	}
	// shift out the incomming address to contain only the data received
	if (n - i + 1) > len(b) {
		copy(b, buf[i+1:i+1+len(b)])
		return n - (i + 1), raddr, errors.New("Datagram did not fit into your buffer.")
	} else {
		copy(b, buf[i+1:n])
		log.WithField("bytesRead", n-(i+1)).Debug("Datagram read successfully")
		return n - (i + 1), raddr, nil
	}
}

func (s *DatagramSession) Accept() (net.Conn, error) {
	log.Debug("Accept called on DatagramSession")
	return s, nil
}

func (s *DatagramSession) Read(b []byte) (n int, err error) {
	log.Debug("Reading from DatagramSession")
	rint, _, rerr := s.ReadFrom(b)
	return rint, rerr
}

// Sends one signed datagram to the destination specified. At the time of
// writing, maximum size is 31 kilobyte, but this may change in the future.
// Implements net.PacketConn.
func (s *DatagramSession) WriteTo(b []byte, addr net.Addr) (n int, err error) {
	log.WithFields(logrus.Fields{
		"addr":        addr,
		"datagramLen": len(b),
	}).Debug("Writing datagram")
	header := []byte("3.1 " + s.ID() + " " + addr.String() + "\n")
	msg := append(header, b...)
	n, err = s.UDPConn.WriteToUDP(msg, s.SAMUDPAddress)
	if err != nil {
		log.WithError(err).Error("Failed to write to UDP")
	} else {
		log.WithField("bytesWritten", n).Debug("Datagram written successfully")
	}
	return n, err
}

func (s *DatagramSession) Write(b []byte) (int, error) {
	log.WithField("dataLen", len(b)).Debug("Writing to DatagramSession")
	return s.WriteTo(b, s.RemoteI2PAddr)
}

// Closes the DatagramSession. Implements net.PacketConn
func (s *DatagramSession) Close() error {
	log.Debug("Closing DatagramSession")
	err := s.Conn.Close()
	err2 := s.UDPConn.Close()
	if err != nil {
		log.WithError(err).Error("Failed to close connection")
		return err
	}
	if err2 != nil {
		log.WithError(err2).Error("Failed to close UDP connection")
	}
	return err2
}

// Returns the I2P destination of the DatagramSession.
func (s *DatagramSession) LocalI2PAddr() i2pkeys.I2PAddr {
	addr := s.DestinationKeys.Addr()
	log.WithField("localI2PAddr", addr).Debug("Getting local I2P address")
	return addr
}

// Implements net.PacketConn
func (s *DatagramSession) LocalAddr() net.Addr {
	return s.LocalI2PAddr()
}

func (s *DatagramSession) Addr() net.Addr {
	return s.LocalI2PAddr()
}

func (s *DatagramSession) Lookup(name string) (a net.Addr, err error) {
	log.WithField("name", name).Debug("Looking up address")
	var sam *common.SAM
	sam, err = common.NewSAM(s.Sam())
	if err == nil {
		defer sam.Close()
		a, err = sam.Lookup(name)
	}
	log.WithField("address", a).Debug("Lookup successful")
	return
}

// Sets read and write deadlines for the DatagramSession. Implements
// net.PacketConn and does the same thing. Setting write deadlines for datagrams
// is seldom done.
func (s *DatagramSession) SetDeadline(t time.Time) error {
	log.WithField("deadline", t).Debug("Setting deadline")
	return s.UDPConn.SetDeadline(t)
}

// Sets read deadline for the DatagramSession. Implements net.PacketConn
func (s *DatagramSession) SetReadDeadline(t time.Time) error {
	log.WithField("readDeadline", t).Debug("Setting read deadline")
	return s.UDPConn.SetReadDeadline(t)
}

// Sets the write deadline for the DatagramSession. Implements net.Packetconn.
func (s *DatagramSession) SetWriteDeadline(t time.Time) error {
	log.WithField("writeDeadline", t).Debug("Setting write deadline")
	return s.UDPConn.SetWriteDeadline(t)
}

func (s *DatagramSession) SetWriteBuffer(bytes int) error {
	log.WithField("bytes", bytes).Debug("Setting write buffer")
	return s.UDPConn.SetWriteBuffer(bytes)
}
