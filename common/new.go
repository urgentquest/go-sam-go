package common

import (
	"fmt"
	"net"
	"strings"
)

// Creates a new controller for the I2P routers SAM bridge.
func OldNewSAM(address string) (*SAM, error) {
	log.WithField("address", address).Debug("Creating new SAM instance")
	var s SAM
	// TODO: clean this up by refactoring the connection setup and error handling logic
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.WithError(err).Error("Failed to dial SAM address")
		return nil, fmt.Errorf("error dialing to address '%s': %w", address, err)
	}
	if _, err := conn.Write(s.SAMEmit.HelloBytes()); err != nil {
		log.WithError(err).Error("Failed to write hello message")
		conn.Close()
		return nil, fmt.Errorf("error writing to address '%s': %w", address, err)
	}
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil {
		log.WithError(err).Error("Failed to read SAM response")
		conn.Close()
		return nil, fmt.Errorf("error reading onto buffer: %w", err)
	}
	if strings.Contains(string(buf[:n]), HELLO_REPLY_OK) {
		log.Debug("SAM hello successful")
		s.SAMEmit.I2PConfig.SetSAMAddress(address)
		s.Conn = conn
		s.SAMResolver, err = NewSAMResolver(&s)
		if err != nil {
			log.WithError(err).Error("Failed to create SAM resolver")
			return nil, fmt.Errorf("error creating resolver: %w", err)
		}
		return &s, nil
	} else if string(buf[:n]) == HELLO_REPLY_NOVERSION {
		log.Error("SAM bridge does not support SAMv3")
		conn.Close()
		return nil, fmt.Errorf("That SAM bridge does not support SAMv3.")
	} else {
		log.WithField("response", string(buf[:n])).Error("Unexpected SAM response")
		conn.Close()
		return nil, fmt.Errorf("%s", string(buf[:n]))
	}
}

func NewSAM(address string) (*SAM, error) {
	logger := log.WithField("address", address)
	logger.Debug("Creating new SAM instance")

	conn, err := connectToSAM(address)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			conn.Close()
		}
	}()

	s := &SAM{
		Conn: conn,
	}

	if err = sendHelloAndValidate(conn, s); err != nil {
		return nil, err
	}

	s.SAMEmit.I2PConfig.SetSAMAddress(address)

	if s.SAMResolver, err = NewSAMResolver(s); err != nil {
		return nil, fmt.Errorf("failed to create SAM resolver: %w", err)
	}

	return s, nil
}
