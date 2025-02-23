package sam3

type SAMResolver struct {
	*SAM
}

func NewSAMResolver(parent *SAM) (*SAMResolver, error) {
	log.Debug("Creating new SAMResolver from existing SAM instance")
	var s SAMResolver
	s.SAM = parent
	return &s, nil
}

func NewFullSAMResolver(address string) (*SAMResolver, error) {
	log.WithField("address", address).Debug("Creating new full SAMResolver")
	var s SAMResolver
	var err error
	s.SAM, err = NewSAM(address)
	if err != nil {
		log.WithError(err).Error("Failed to create new SAM instance")
		return nil, err
	}
	return &s, nil
}
