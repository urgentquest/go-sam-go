package sam3

import (
	"github.com/go-i2p/go-sam-go/stream"
)

// Implements net.Conn
type SAMConn struct {
	*stream.StreamConn
}
