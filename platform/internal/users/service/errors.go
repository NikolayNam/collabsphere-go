package service

import "errors"

var (
	ErrValidation = errors.New("validation")
	ErrConflict   = errors.New("conflict")
	ErrInternal   = errors.New("internal")
)
