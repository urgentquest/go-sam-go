package datagram

import (
	"net"

	"github.com/go-i2p/go-sam-go/common"
	"github.com/go-i2p/i2pkeys"
)

type SAM common.SAM

// The DatagramSession implements net.PacketConn. It works almost like ordinary
// UDP, except that datagrams may be at most 31kB large. These datagrams are
// also end-to-end encrypted, signed and includes replay-protection. And they
// are also built to be surveillance-resistant (yey!).
type DatagramSession struct {
	*SAM
	UDPConn       *net.UDPConn     // used to deliver datagrams
	SAMUDPAddress *net.UDPAddr     // the SAM bridge UDP-port
	RemoteI2PAddr *i2pkeys.I2PAddr // optional remote I2P address
}
