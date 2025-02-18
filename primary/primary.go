package primary

/*
// Creates a new PrimarySession with the I2CP- and streaminglib options as
// specified. See the I2P documentation for a full list of options.
func (sam *SAM) NewPrimarySession(id string, keys i2pkeys.I2PKeys, options []string) (*PrimarySession, error) {
	log.WithFields(logrus.Fields{"id": id, "options": options}).Debug("NewPrimarySession() called")
	return sam.newPrimarySession(PrimarySessionSwitch, id, keys, options)
}

func (sam *SAM) newPrimarySession(primarySessionSwitch string, id string, keys i2pkeys.I2PKeys, options []string) (*PrimarySession, error) {
	log.WithFields(logrus.Fields{
		"primarySessionSwitch": primarySessionSwitch,
		"id":                   id,
		"options":              options,
	}).Debug("newPrimarySession() called")

	conn, err := sam.newGenericSession(primarySessionSwitch, id, keys, options, []string{})
	if err != nil {
		log.WithError(err).Error("Failed to create new generic session")
		return nil, err
	}
	ssesss := make(map[string]*StreamSession)
	dsesss := make(map[string]*DatagramSession)
	return &PrimarySession{sam.Config.I2PConfig.Sam(), id, conn, keys, time.Duration(600 * time.Second), time.Now(), Sig_NONE, sam.Config, ssesss, dsesss}, nil
}

// Creates a new PrimarySession with the I2CP- and PRIMARYinglib options as
// specified. See the I2P documentation for a full list of options.
func (sam *SAM) NewPrimarySessionWithSignature(id string, keys i2pkeys.I2PKeys, options []string, sigType string) (*PrimarySession, error) {
	log.WithFields(logrus.Fields{
		"id":      id,
		"options": options,
		"sigType": sigType,
	}).Debug("NewPrimarySessionWithSignature() called")

	conn, err := sam.newGenericSessionWithSignature(PrimarySessionSwitch, id, keys, sigType, options, []string{})
	if err != nil {
		log.WithError(err).Error("Failed to create new generic session with signature")
		return nil, err
	}
	ssesss := make(map[string]*stream.StreamSession)
	dsesss := make(map[string]*datagram.DatagramSession)
	return &PrimarySession{sam.Config.I2PConfig.Sam(), id, conn, keys, time.Duration(600 * time.Second), time.Now(), sigType, sam.Config, ssesss, dsesss}, nil
}
*/
