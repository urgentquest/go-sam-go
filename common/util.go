package common

import (
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func IgnorePortError(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "missing port in address") {
		log.Debug("Ignoring 'missing port in address' error")
		err = nil
	}
	return err
}

func SplitHostPort(hostport string) (string, string, error) {
	host, port, err := net.SplitHostPort(hostport)
	if err != nil {
		if IgnorePortError(err) == nil {
			log.WithField("host", hostport).Debug("Using full string as host, port set to 0")
			host = hostport
			port = "0"
		}
	}
	log.WithFields(logrus.Fields{
		"host": host,
		"port": port,
	}).Debug("Split host and port")
	return host, port, nil
}

var randSource = rand.NewSource(time.Now().UnixNano())
var randGen = rand.New(randSource)

func RandPort() string {
	for {
		p := randGen.Intn(55534) + 10000
		port := strconv.Itoa(p)
		if l, e := net.Listen("tcp", net.JoinHostPort("localhost", port)); e != nil {
			continue
		} else {
			defer l.Close()
			if l, e := net.Listen("udp", net.JoinHostPort("localhost", port)); e != nil {
				continue
			} else {
				defer l.Close()
				return strconv.Itoa(l.Addr().(*net.UDPAddr).Port)
			}
		}
	}
}
