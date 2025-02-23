// package sam3 wraps the original sam3 API from github.com/go-i2p/sam3
package sam3

import (
	"github.com/go-i2p/go-sam-go/datagram"
)

// The DatagramSession implements net.PacketConn. It works almost like ordinary
// UDP, except that datagrams may be at most 31kB large. These datagrams are
// also end-to-end encrypted, signed and includes replay-protection. And they
// are also built to be surveillance-resistant (yey!).
type DatagramSession struct {
	datagram.DatagramSession
}
