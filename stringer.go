package stringx

import (
	"fmt"
)

type Stringer interface {
	String() string
}

func NewStringer(s Stringer) Stringer {
	return stringer{Stringer: s}
}

type stringer struct {
	Stringer
}

func (w stringer) String() string {
	return fmt.Sprintf("%v", w.Stringer)
}
