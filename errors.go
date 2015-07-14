package tgbot

import (
	"errors"
)

var (
	ErrorInvalidArgs    error = errors.New("Invalid Arguments")
	ErrorNotImplemented       = errors.New("Not Implemented")
)
