package common

import (
	"fmt"
	"net"
	"strings"
)

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

func connectToSAM(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SAM bridge at %s: %w", address, err)
	}
	return conn, nil
}

func sendHelloAndValidate(conn net.Conn, s *SAM) error {
	if _, err := conn.Write(s.SAMEmit.HelloBytes()); err != nil {
		return fmt.Errorf("failed to send hello message: %w", err)
	}

	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("failed to read SAM response: %w", err)
	}

	response := string(buf[:n])
	switch {
	case strings.Contains(response, HELLO_REPLY_OK):
		log.Debug("SAM hello successful")
		return nil
	case response == HELLO_REPLY_NOVERSION:
		return fmt.Errorf("SAM bridge does not support SAMv3")
	default:
		return fmt.Errorf("unexpected SAM response: %s", response)
	}
}
