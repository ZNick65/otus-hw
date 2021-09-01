package main

import (
	"testing"

	"github.com/go-playground/assert"
)

var tests = []struct {
	name         string
	greeting     string
	expectations string
	valid        bool
}{
	{
		"valid",
		"Hello, OTUS!",
		"!SUTO ,olleH",
		true,
	},
	{
		"no match",
		"Hello, OTUS!",
		"SUTO ,olleH",
		false,
	},
	{
		"empty",
		"",
		"",
		true,
	},
}

func TestReverseString(t *testing.T) {
	for _, test := range tests {
		greeting := reverseString(test.greeting)
		if test.valid {
			assert.Equal(t, test.expectations, greeting)
		} else {
			assert.NotEqual(t, test.expectations, greeting)
		}
	}
}
