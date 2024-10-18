package domain

import "errors"

var (
	ErrIncorrectConfirmPassword = errors.New("incorrect confirm password")
	ErrWrongPassword            = errors.New("wrong password")
	ErrUnableToDeleteUser       = errors.New("unable to delete user")
	ErrUnexpectedSigningMethod  = errors.New("unexpected signing method")
	ErrInvalidTokenClaims       = errors.New("invalid token claims")
	ErrMissingParams            = errors.New("missing params")
)
