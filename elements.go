package elements

import "errors"

var (
	// ErrAttachNotSupported returned when attaching elments is not supported
	ErrAttachNotSupported = errors.New("attach not supported")
)
