package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(unpack string) (string, error) {
	mup := NewMyUnpack()
	return mup.Unpack(unpack)
}
