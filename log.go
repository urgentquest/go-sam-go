// package sam3 wraps the original sam3 API from github.com/go-i2p/sam3
package sam3

import (
	"io"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	log  *logrus.Logger
	once sync.Once
)

func InitializeSAM3Logger() {
	once.Do(func() {
		log = logrus.New()
		// We do not want to log by default
		log.SetOutput(io.Discard)
		log.SetLevel(logrus.PanicLevel)
		// Check if DEBUG_I2P is set
		if logLevel := os.Getenv("DEBUG_I2P"); logLevel != "" {
			log.SetOutput(os.Stdout)
			switch strings.ToLower(logLevel) {
			case "debug":
				log.SetLevel(logrus.DebugLevel)
			case "warn":
				log.SetLevel(logrus.WarnLevel)
			case "error":
				log.SetLevel(logrus.ErrorLevel)
			default:
				log.SetLevel(logrus.DebugLevel)
			}
			log.WithField("level", log.GetLevel()).Debug("Logging enabled.")
		}
	})
}

// GetSAM3Logger returns the initialized logger
func GetSAM3Logger() *logrus.Logger {
	if log == nil {
		InitializeSAM3Logger()
	}
	return log
}

func init() {
	InitializeSAM3Logger()
}
