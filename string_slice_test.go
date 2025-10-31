package stringx_test

import (
	"github.com/franela/goblin"
	"github.com/noho-digital/insurews/pkg/suite"
	"github.com/noho-digital/insurews/pkg/test/reporter"
	"testing"

	. "github.com/noho-digital/insurews/pkg/strings"
	. "github.com/onsi/gomega"
)

func TestStringSlice(t *testing.T) {
	suite.Run(t, new(StringSetSuite))
}

type StringSliceSuite struct {
	suite.Suite
}

func (s *StringSliceSuite) TestAnyString() {
	g := goblin.Goblin(s.T())
	g.SetReporter(reporter.NewReporter(reporter.WithDefaultsFromSettings(s.Settings)))
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	g.Describe("String  slice", func() {
		g.It("Can use AnyEqual to find contents", func() {
			sl := StringSlice{"a", "b", "c", "d"}
			for _, s := range sl {
				Expect(sl.AnyEqual(s)).To(BeTrue())
			}
			Expect(sl.AnyEqual("z")).NotTo(BeTrue())
		})
	})
}
