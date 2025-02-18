package stream

import (
	"net"
	"time"

	"github.com/go-i2p/i2pkeys"
)

// Implements net.Conn
func (sc *StreamConn) Read(buf []byte) (int, error) {
	n, err := sc.conn.Read(buf)
	return n, err
}

// Implements net.Conn
func (sc *StreamConn) Write(buf []byte) (int, error) {
	n, err := sc.conn.Write(buf)
	return n, err
}

// Implements net.Conn
func (sc *StreamConn) Close() error {
	return sc.conn.Close()
}

func (sc *StreamConn) LocalAddr() net.Addr {
	return sc.localAddr()
}

// Implements net.Conn
func (sc *StreamConn) localAddr() i2pkeys.I2PAddr {
	return sc.laddr
}

func (sc *StreamConn) RemoteAddr() net.Addr {
	return sc.remoteAddr()
}

// Implements net.Conn
func (sc *StreamConn) remoteAddr() i2pkeys.I2PAddr {
	return sc.raddr
}

// Implements net.Conn
func (sc *StreamConn) SetDeadline(t time.Time) error {
	return sc.conn.SetDeadline(t)
}

// Implements net.Conn
func (sc *StreamConn) SetReadDeadline(t time.Time) error {
	return sc.conn.SetReadDeadline(t)
}

// Implements net.Conn
func (sc *StreamConn) SetWriteDeadline(t time.Time) error {
	return sc.conn.SetWriteDeadline(t)
}
