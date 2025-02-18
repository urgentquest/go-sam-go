package primary

import (
	"net"
	"time"

	"github.com/go-i2p/go-sam-go/common"
	"github.com/go-i2p/go-sam-go/datagram"
	"github.com/go-i2p/go-sam-go/stream"
	"github.com/go-i2p/i2pkeys"
)

type SAM common.SAM

// Represents a primary session.
type PrimarySession struct {
	*SAM
	samAddr  string          // address to the sam bridge (ipv4:port)
	id       string          // tunnel name
	conn     net.Conn        // connection to sam
	keys     i2pkeys.I2PKeys // i2p destination keys
	Timeout  time.Duration
	Deadline time.Time
	sigType  string
	Config   common.SAMEmit
	stsess   map[string]*stream.StreamSession
	dgsess   map[string]*datagram.DatagramSession
	//	from     string
	//	to       string
}
