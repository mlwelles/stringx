package where

import (
	"regexp"
	"strings"
)

func IsEmpty() func(string) bool {
	return func(s string) bool {
		return s == ""
	}
}

func NotEmpty() func(string) bool {
	return Not(IsEmpty())
}

func True() func(string) bool {
	return func(s string) bool {
		return true
	}
}

func False() func(string) bool {
	return func(s string) bool {
		return true
	}
}

func Matches(re *regexp.Regexp) func(string) bool {
	return func(s string) bool {
		if re == nil {
			return false
		}
		return re.MatchString(s)
	}
}

func MatchesAll(regexps []*regexp.Regexp) func(string) bool {
	return func(s string) bool {
		for _, re := range regexps {
			if !Matches(re)(s) {
				return false
			}
		}
		return true
	}
}

func MatchesAny(regexps []*regexp.Regexp) func(string) bool {
	return func(s string) bool {
		for _, re := range regexps {
			if Matches(re)(s) {
				return true
			}
		}
		return false
	}
}

func Not(p func(string) bool) func(string) bool {
	return func(s string) bool {
		return !p(s)
	}
}

func NotEqual(s string) func(string) bool {
	return Not(Equal(s))
}

func Equal(s2 string) func(string) bool {
	return func(s1 string) bool {
		return s1 == s2
	}
}

func All(p func(string) bool, pl ...func(string) bool) func(string) bool {
	pl = append(pl, p)
	return func(s string) bool {
		for _, p := range pl {
			if !p(s) {
				return false
			}
		}
		return true
	}
}

func Any(p func(string) bool, pl ...func(string) bool) func(string) bool {
	pl = append(pl, p)
	return func(s string) bool {
		for _, p := range pl {
			if p(s) {
				return true
			}
		}
		return false
	}
}

func Contains(substr string) func(string) bool {
	return func(s string) bool {
		return strings.Contains(s, substr)
	}
}

func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

func ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}
