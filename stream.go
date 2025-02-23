// package sam3 wraps the original sam3 API from github.com/go-i2p/sam3
package sam3

import (
	"time"

	"github.com/go-i2p/go-sam-go/stream"
)

// Represents a streaming session.
type StreamSession struct {
	*stream.StreamSession
}

/*
func (s *StreamSession) Cancel() chan *StreamSession {
	ch := make(chan *StreamSession)
	ch <- s
	return ch
}*/

func minNonzeroTime(a, b time.Time) time.Time {
	if a.IsZero() {
		return b
	}
	if b.IsZero() || a.Before(b) {
		return a
	}
	return b
}
