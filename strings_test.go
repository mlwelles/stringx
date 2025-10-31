package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanAlpha(t *testing.T) {
	a := assert.New(t)
	es := StringSlice{
		"",
		"#@@",
		"-@#!",
	}
	for _, i := range es {
		o := CleanAlphaNumeric(i)
		a.Emptyf(o, "%q after clean alpha should be empty string, got %q", i, o)
	}

	es = StringSlice{
		" ",
		"a-d112@#32q",
		" a",
		"- ",
	}

	for _, i := range es {
		o := CleanAlphaNumeric(i)
		a.NotEmptyf(o, "%q after clean alpha should be empty string, got %q", i, o)
	}
}
func TestStringsEmptyByteArray(t *testing.T) {
	a := assert.New(t)
	s := string([]byte(""))
	a.Emptyf(s, "%q should be empty", s)
}
