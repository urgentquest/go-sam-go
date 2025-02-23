// package sam3 wraps the original sam3 API from github.com/go-i2p/sam3
package sam3

import (
	"github.com/go-i2p/go-sam-go/stream"
)

type StreamListener struct {
	*stream.StreamListener
}
