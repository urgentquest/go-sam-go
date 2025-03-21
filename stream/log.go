package stream

import logger "github.com/go-i2p/go-sam-go/logger"

var log = logger.GetSAM3Logger()

func init() {
	logger.InitializeSAM3Logger()
	log = logger.GetSAM3Logger()
}
