package stringx

import "strings"

type Replacer = strings.Replacer

// NewReplacer returns a new Replacer from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
//
// NewReplacer panics if given an odd number of arguments.
func NewReplacer(oldnew ...string) *Replacer {
	return strings.NewReplacer(oldnew...)
}
