package raw

import (
	"errors"
	"net"
	"strconv"

	"github.com/go-i2p/go-sam-go/common"
	"github.com/go-i2p/i2pkeys"
	"github.com/sirupsen/logrus"
)

// Creates a new raw session. udpPort is the UDP port SAM is listening on,
// and if you set it to zero, it will use SAMs standard UDP port.
func (s *SAM) NewRawSession(id string, keys i2pkeys.I2PKeys, options []string, udpPort int) (*RawSession, error) {
	log.WithFields(logrus.Fields{"id": id, "udpPort": udpPort}).Debug("Creating new RawSession")

	if udpPort > 65335 || udpPort < 0 {
		log.WithField("udpPort", udpPort).Error("Invalid UDP port")
		return nil, errors.New("udpPort needs to be in the interval 0-65335")
	}
	if udpPort == 0 {
		udpPort = 7655
		log.Debug("Using default UDP port 7655")
	}
	lhost, _, err := common.SplitHostPort(s.LocalAddr().String())
	if err != nil {
		log.Debug("Using default UDP port 7655")
		s.Close()
		return nil, err
	}
	lUDPAddr, err := net.ResolveUDPAddr("udp4", lhost+":0")
	if err != nil {
		log.WithError(err).Error("Failed to resolve local UDP address")
		return nil, err
	}
	udpconn, err := net.ListenUDP("udp4", lUDPAddr)
	if err != nil {
		log.WithError(err).Error("Failed to listen on UDP")
		return nil, err
	}
	rhost, _, err := common.SplitHostPort(s.RemoteAddr().String())
	if err != nil {
		log.WithError(err).Error("Failed to split remote host port")
		s.Close()
		return nil, err
	}
	rUDPAddr, err := net.ResolveUDPAddr("udp4", rhost+":"+strconv.Itoa(udpPort))
	if err != nil {
		log.WithError(err).Error("Failed to resolve remote UDP address")
		return nil, err
	}
	_, lport, err := net.SplitHostPort(udpconn.LocalAddr().String())
	if err != nil {
		log.WithError(err).Error("Failed to get local port")
		return nil, err
	}
	conn, err := s.NewGenericSession("RAW", id, keys, []string{"PORT=" + lport})
	if err != nil {
		log.WithError(err).Error("Failed to create new generic session")
		return nil, err
	}
	log.WithFields(logrus.Fields{
		"id":            id,
		"localPort":     lport,
		"remoteUDPAddr": rUDPAddr,
	}).Debug("Created new RawSession")
	rawSession := &RawSession{
		SAM: s,
	}
	rawSession.Conn = conn
	return rawSession, nil
}
