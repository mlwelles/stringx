package stringx_test

import (
	"github.com/mlwelles/suite"
	"testing"
)

type StringSetSuite struct {
	suite.Suite
}

func TestStringSet(t *testing.T) {
	suite.Run(t, new(StringSetSuite))
}

// FIXME -- re-enable and fix the generated set iterator
//func (s *StringSetSuite) TestStringSet() {
//	g := Goblin(s.T())
//	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
//	contents := StringSlice{"a", "b", "c", "d", "e"}
//	set := NewStringSet(contents...)
//
//	g.Describe("Generated string set works", func() {
//		g.It("Can iterate over members", func() {
//			iter := set.Iterator()
//			Expect(iter).ToNot(BeNil(), "iterator should not be nil")
//			index := 0
//			for str := range iter.C {
//				hasAny := contents.Any(func(s string) bool { return s == str })
//				Expect(hasAny).To(BeTrue(), fmt.Sprintf("item %d returned by iterator, %s should be in %s", index, str, contents))
//			}
//		})
//	})
// }
