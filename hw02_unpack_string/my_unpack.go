package hw02unpackstring

import (
	"strconv"
	"strings"
	"unicode"
)

// MyUnpack - the struct for unpacking srting.
type MyUnpack struct {
	sb        strings.Builder
	backslash bool
	last      rune
}

// NewMyUnpack - create the unpacking struct.
func NewMyUnpack() *MyUnpack {
	return &MyUnpack{
		sb:        strings.Builder{},
		backslash: false,
		last:      0,
	}
}

// Unpack - main function.
func (u *MyUnpack) Unpack(_s string) (string, error) {
	if !u.valid(_s) {
		return "", ErrInvalidString
	}

	for _, r := range _s {
		switch {
		case u.isBackslash(r):
			u.addBackslash(r)
		case u.isDigit(r) && !u.backslash:
			if err := u.repeatLast(r); err != nil {
				return "", err
			}
		case u.isDigit(r) && u.backslash:
			u.add(r)
			u.backslash = false
		default:
			u.add(r)
		}
	}
	return u.string(), nil
}

// valid - check a string is valid.
func (u *MyUnpack) valid(s string) bool {
	var prev rune
	for i, r := range s {
		switch {
		case i == 0 && u.isDigit(r):
			return false
		case i > 0 && u.isDigit(r) && u.isDigit(prev):
			return false
		case i > 0 && u.isBackslash(prev) && !(u.isDigit(r) || u.isBackslash(r)):
			return false
		}

		if !u.isBackslash(prev) {
			prev = r
		}
	}
	return true
}

// addBackslash ...
// if backslash with a shild then fixing one, if first time set shild.
func (u *MyUnpack) addBackslash(r rune) {
	if u.backslash {
		u.add(r)
		u.backslash = false
	} else {
		u.backslash = true
	}
}

// isDigit - checking a rune is digit.
func (u *MyUnpack) isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// isBackslash - checking a rune is backslash.
func (u *MyUnpack) isBackslash(r rune) bool {
	return string(r) == "\\"
}

// string - get a current string.
func (u *MyUnpack) string() string {
	return u.sb.String()
}

// add - add rune to builder.
func (u *MyUnpack) add(r rune) {
	u.last = r
	u.sb.WriteRune(r)
}

// str - add string to builder.
func (u *MyUnpack) str(s string) {
	u.sb.WriteString(s)
}

// len - get current len of inside buffer.
func (u *MyUnpack) len() int {
	return u.sb.Len()
}

// strFirst - return current string without n last runes.
func (u *MyUnpack) strFirst(n int) (s string) {
	if u.len()+n != 0 {
		s = (u.string())[:u.len()+n]
	}
	return
}

// repeat a rune n times.
func (u *MyUnpack) repeat(r rune, n int) {
	if n > 0 {
		u.str(strings.Repeat(string(r), n))
	}
}

// reset - clean the inside buffer.
func (u *MyUnpack) reset() {
	u.sb.Reset()
}

// repeatLast - repeat the last rune n times.
func (u *MyUnpack) repeatLast(r rune) error {
	if u.len() != 0 {
		d, er := strconv.Atoi(string(r))
		if er != nil {
			return ErrInvalidString
		}

		first := u.strFirst(-1)
		u.reset()
		u.str(first)
		u.repeat(u.last, d)
		u.last = r
	}
	return nil
}
