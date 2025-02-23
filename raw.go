package sam3

import (
	"github.com/go-i2p/go-sam-go/raw"
)

// The RawSession provides no authentication of senders, and there is no sender
// address attached to datagrams, so all communication is anonymous. The
// messages send are however still endpoint-to-endpoint encrypted. You
// need to figure out a way to identify and authenticate clients yourself, iff
// that is needed. Raw datagrams may be at most 32 kB in size. There is no
// overhead of authentication, which is the reason to use this..
type RawSession struct {
	*raw.RawSession
}
