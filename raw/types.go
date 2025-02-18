package raw

import (
	"net"

	"github.com/go-i2p/go-sam-go/common"
)

type SAM common.SAM

// The RawSession provides no authentication of senders, and there is no sender
// address attached to datagrams, so all communication is anonymous. The
// messages send are however still endpoint-to-endpoint encrypted. You
// need to figure out a way to identify and authenticate clients yourself, iff
// that is needed. Raw datagrams may be at most 32 kB in size. There is no
// overhead of authentication, which is the reason to use this..
type RawSession struct {
	*SAM
	SAMUDPConn *net.UDPConn // used to deliver datagrams
	SAMUDPAddr *net.UDPAddr // the SAM bridge UDP-port
}
