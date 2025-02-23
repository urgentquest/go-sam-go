// package sam3 wraps the original sam3 API from github.com/go-i2p/sam3
package sam3

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/go-i2p/go-sam-go/primary"
)

const (
	session_ADDOK = "SESSION STATUS RESULT=OK"
)

func randport() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	p := r.Intn(55534) + 10000
	port := strconv.Itoa(p)
	log.WithField("port", port).Debug("Generated random port")
	return strconv.Itoa(p)
}

// Represents a primary session.
type PrimarySession struct {
	*primary.PrimarySession
}

var PrimarySessionSwitch = "MASTER"

func (p *PrimarySession) NewStreamSubSession(id string) (*StreamSession, error) {
	log.WithField("id", id).Debug("NewStreamSubSession called")
	session, err := p.PrimarySession.NewStreamSubSession(id)
	if err != nil {
		return nil, err
	}
	return &StreamSession{
		StreamSession: session,
	}, nil
}
