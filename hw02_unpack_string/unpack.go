package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(unpack string) (string, error) {
	// Place your code here.
	return NewMyUnpack().Unpack(unpack)
}
