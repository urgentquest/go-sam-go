package primary

import (
	"errors"
	"net"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/go-i2p/go-sam-go/common"
)

// Creates a new session with the style of either "STREAM", "DATAGRAM" or "RAW",
// for a new I2P tunnel with name id, using the cypher keys specified, with the
// I2CP/streaminglib-options as specified. Extra arguments can be specified by
// setting extra to something else than []string{}.
// This sam3 instance is now a session
func (sam *PrimarySession) NewGenericSubSession(style, id string, extras []string) (net.Conn, error) {
	log.WithFields(logrus.Fields{"style": style, "id": id, "extras": extras}).Debug("newGenericSubSession called")
	return sam.NewGenericSubSessionWithSignature(style, id, extras)
}

func (sam *PrimarySession) NewGenericSubSessionWithSignature(style, id string, extras []string) (net.Conn, error) {
	log.WithFields(logrus.Fields{"style": style, "id": id, "extras": extras}).Debug("newGenericSubSessionWithSignature called")
	return sam.NewGenericSubSessionWithSignatureAndPorts(style, id, "0", "0", extras)
}

// Creates a new session with the style of either "STREAM", "DATAGRAM" or "RAW",
// for a new I2P tunnel with name id, using the cypher keys specified, with the
// I2CP/streaminglib-options as specified. Extra arguments can be specified by
// setting extra to something else than []string{}.
// This sam3 instance is now a session
func (sam *PrimarySession) NewGenericSubSessionWithSignatureAndPorts(style, id, from, to string, extras []string) (net.Conn, error) {
	log.WithFields(logrus.Fields{"style": style, "id": id, "from": from, "to": to, "extras": extras}).Debug("newGenericSubSessionWithSignatureAndPorts called")

	conn := sam.conn
	fp := ""
	tp := ""
	if from != "0" && from != "" {
		fp = " FROM_PORT=" + from
	}
	if to != "0" && to != "" {
		tp = " TO_PORT=" + to
	}
	scmsg := []byte("SESSION ADD STYLE=" + style + " ID=" + id + fp + tp + " " + strings.Join(extras, " ") + "\n")

	log.WithField("message", string(scmsg)).Debug("Sending SESSION ADD message")

	for m, i := 0, 0; m != len(scmsg); i++ {
		if i == 15 {
			conn.Close()
			log.Error("Writing to SAM failed after 15 attempts")
			return nil, errors.New("writing to SAM failed")
		}
		n, err := conn.Write(scmsg[m:])
		if err != nil {
			log.WithError(err).Error("Failed to write to SAM connection")
			conn.Close()
			return nil, err
		}
		m += n
	}
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		log.WithError(err).Error("Failed to read from SAM connection")
		conn.Close()
		return nil, err
	}
	text := string(buf[:n])
	log.WithField("response", text).Debug("Received response from SAM")
	// log.Println("SAM:", text)
	if strings.HasPrefix(text, SESSION_ADDOK) {
		//if sam.keys.String() != text[len(common.SESSION_ADDOK):len(text)-1] {
		//conn.Close()
		//return nil, errors.New("SAMv3 created a tunnel with keys other than the ones we asked it for")
		//}
		log.Debug("Session added successfully")
		return conn, nil //&StreamSession{id, conn, keys, nil, sync.RWMutex{}, nil}, nil
	} else if text == common.SESSION_DUPLICATE_ID {
		log.Error("Duplicate tunnel name")
		conn.Close()
		return nil, errors.New("Duplicate tunnel name")
	} else if text == common.SESSION_DUPLICATE_DEST {
		log.Error("Duplicate destination")
		conn.Close()
		return nil, errors.New("Duplicate destination")
	} else if text == common.SESSION_INVALID_KEY {
		log.Error("Invalid key - Primary Session")
		conn.Close()
		return nil, errors.New("Invalid key - Primary Session")
	} else if strings.HasPrefix(text, common.SESSION_I2P_ERROR) {
		log.WithField("error", text[len(common.SESSION_I2P_ERROR):]).Error("I2P error")
		conn.Close()
		return nil, errors.New("I2P error " + text[len(common.SESSION_I2P_ERROR):])
	} else {
		log.WithField("reply", text).Error("Unable to parse SAMv3 reply")
		conn.Close()
		return nil, errors.New("Unable to parse SAMv3 reply: " + text)
	}
}
