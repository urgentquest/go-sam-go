package primary

import (
	"time"

	"github.com/go-i2p/go-sam-go/common"
	"github.com/go-i2p/go-sam-go/datagram"
	"github.com/go-i2p/go-sam-go/stream"
	"github.com/go-i2p/i2pkeys"
	"github.com/sirupsen/logrus"
)

var PrimarySessionSwitch string = "MASTER"

// Creates a new PrimarySession with the I2CP- and streaminglib options as
// specified. See the I2P documentation for a full list of options.
func (sam *SAM) NewPrimarySession(id string, keys i2pkeys.I2PKeys, options []string) (*PrimarySession, error) {
	log.WithFields(logrus.Fields{"id": id, "options": options}).Debug("NewPrimarySession() called")
	return sam.newPrimarySession(PrimarySessionSwitch, id, keys, options)
}

func (sam *SAM) newPrimarySession(primarySessionSwitch, id string, keys i2pkeys.I2PKeys, options []string) (*PrimarySession, error) {
	log.WithFields(logrus.Fields{
		"primarySessionSwitch": primarySessionSwitch,
		"id":                   id,
		"options":              options,
	}).Debug("newPrimarySession() called")

	conn, err := sam.NewGenericSession(primarySessionSwitch, id, keys, options)
	if err != nil {
		log.WithError(err).Error("Failed to create new generic session")
		return nil, err
	}
	return &PrimarySession{
		SAM:      sam,
		samAddr:  "",
		id:       id,
		conn:     conn,
		keys:     keys,
		Timeout:  0,
		Deadline: time.Time{},
		sigType:  "",
		Config:   common.SAMEmit{},
		stsess:   map[string]*stream.StreamSession{},
		dgsess:   map[string]*datagram.DatagramSession{},
	}, nil
}

// Creates a new PrimarySession with the I2CP- and PRIMARYinglib options as
// specified. See the I2P documentation for a full list of options.
func (sam *SAM) NewPrimarySessionWithSignature(id string, keys i2pkeys.I2PKeys, options []string, sigType string) (*PrimarySession, error) {
	log.WithFields(logrus.Fields{
		"id":      id,
		"options": options,
		"sigType": sigType,
	}).Debug("NewPrimarySessionWithSignature() called")

	conn, err := sam.NewGenericSessionWithSignature(PrimarySessionSwitch, id, keys, sigType, options)
	if err != nil {
		log.WithError(err).Error("Failed to create new generic session with signature")
		return nil, err
	}
	return &PrimarySession{
		SAM:      sam,
		samAddr:  "",
		id:       id,
		conn:     conn,
		keys:     keys,
		Timeout:  0,
		Deadline: time.Time{},
		sigType:  sigType,
		Config:   common.SAMEmit{},
		stsess:   map[string]*stream.StreamSession{},
		dgsess:   map[string]*datagram.DatagramSession{},
	}, nil
}
