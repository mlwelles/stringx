package stringx

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/mlwelles/where"
)

func IsEmpty(s string) bool {
	return where.IsEmpty()(s)
}

func NotEmpty(s string) bool {
	return where.NotEmpty()(s)
}

func Match(s string, re *regexp.Regexp) bool {
	return where.Matches(re)(s)
}

func All(strs []string, p func(string) bool) bool {
	for _, s := range strs {
		if !p(s) {
			return false
		}
	}
	return true
}

func Any(strs []string, p func(string) bool) bool {
	for _, s := range strs {
		if p(s) {
			return true
		}
	}
	return false
}

func AllMatch(strs []string, re *regexp.Regexp) bool {
	return All(strs, where.Matches(re))
}

func MatchAll(s string, regexps []*regexp.Regexp) bool {
	return where.MatchesAll(regexps)(s)
}

func MatchAny(s string, regexps []*regexp.Regexp) bool {
	return where.MatchesAny(regexps)(s)

}

func AnyMatch(strs []string, re *regexp.Regexp) bool {
	return Any(strs, where.Matches(re))
}

func Concat(slice []string, slices ...[]string) StringSlice {
	for _, sl := range slices {
		slice = append(slice, sl...)
	}
	return slice
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte(s string, c byte) int {
	return strings.IndexByte(s, c)
}

// IndexRune returns the index of the first instance of the Unicode code point
// r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any
// invalid UTF-8 byte sequence.
func IndexRune(s string, r rune) int {
	return strings.IndexRune(s, r)
}

// IndexAny returns the index of the first instance of any Unicode code point
// from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// LastIndexAny returns the index of the last instance of any Unicode code
// point from chars in s, or -1 if no Unicode code point from chars is
// present in s.
func LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte(s string, c byte) int {
	return strings.LastIndexByte(s, c)
}

// SplitN slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
func SplitN(s, sep string, n int) StringSlice {
	return strings.SplitN(s, sep, n)
}

// SplitAfterN slices s into substrings after each instance of sep and
// returns a slice of those substrings.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings
func SplitAfterN(s, sep string, n int) StringSlice {
	return strings.SplitAfterN(s, sep, n)
}

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
func Split(s, sep string) StringSlice {
	return strings.Split(s, sep)
}

// SplitAfter slices s into all substrings after each instance of sep and
// returns a slice of those substrings.
//
// If s does not contain sep and sep is not empty, SplitAfter returns
// a slice of length 1 whose only element is s.
//
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If
// both s and sep are empty, SplitAfter returns an empty slice.
//
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter(s, sep string) StringSlice {
	return strings.SplitAfter(s, sep)
}

// Fields splits the string s around each instance of one or more consecutive white space
// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
// empty slice if s contains only white space.
func Fields(s string) StringSlice {
	return strings.Fields(s)
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
// and returns an array of slices of s. If all code points in s satisfy f(c) or the
// string is empty, an empty slice is returned.
// FieldsFunc makes no guarantees about the order in which it calls f(c).
// If f does not return consistent results for a given c, FieldsFunc may crash.
func FieldsFunc(s string, f func(rune) bool) StringSlice {
	return strings.FieldsFunc(s, f)
}

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasAnySuffix tests whether the string ends with any of the  suffixes
func HasAnySuffix(s string, suffixes ...string) bool {
	for _, sfx := range suffixes {
		if HasSuffix(s, sfx) {
			return true
		}
	}
	return false
}

// HasSuffix tests whether the string s ends with suffix
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Map returns a copy of the string s with all its characters modified
// according to the mapping function. If mapping returns a negative value, the character is
// dropped from the string with no replacement.
func Map(mapping func(rune) rune, s string) string {
	return strings.Map(mapping, s)
}

// Repeat returns a new string consisting of count copies of the string s.
//
// It panics if count is negative or if
// the result of (len(s) * count) overflows.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
func ToTitle(s string) string {
	return strings.ToTitle(s)
}

// ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their
// upper case using the case mapping specified by c.
func ToUpperSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToUpperSpecial(c, s)
}

// ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their
// lower case using the case mapping specified by c.
func ToLowerSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToLowerSpecial(c, s)
}

// ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their
// Unicode title case, giving priority to the special casing rules.
func ToTitleSpecial(c unicode.SpecialCase, s string) string {
	return strings.ToTitleSpecial(c, s)
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
func ToValidUTF8(s, replacement string) string {
	return strings.ToValidUTF8(s, replacement)
}

// Title returns a copy of the string s with all Unicode letters that begin words
// mapped to their Unicode title case.
//
// BUG(rsc): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
func Title(s string) string {
	return strings.Title(s)
}

// TrimLeftFunc returns a slice of the string s with all leading
// Unicode code points c satisfying f(c) removed.
func TrimLeftFunc(s string, f func(rune) bool) string {
	return strings.TrimLeftFunc(s, f)
}

// TrimRightFunc returns a slice of the string s with all trailing
// Unicode code points c satisfying f(c) removed.
func TrimRightFunc(s string, f func(rune) bool) string {
	return strings.TrimRightFunc(s, f)
}

// TrimFunc returns a slice of the string s with all leading
// and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc(s string, f func(rune) bool) string {
	return strings.TrimFunc(s, f)
}

// IndexFunc returns the index into s of the first Unicode
// code point satisfying f(c), or -1 if none do.
func IndexFunc(s string, f func(rune) bool) int {
	return strings.IndexFunc(s, f)
}

// LastIndexFunc returns the index into s of the last
// Unicode code point satisfying f(c), or -1 if none do.
func LastIndexFunc(s string, f func(rune) bool) int {
	return strings.LastIndexFunc(s, f)
}

// Trim returns a slice of the string s with all leading and
// trailing Unicode code points contained in cutset removed.
func Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

// TrimLeft returns a slice of the string s with all leading
// Unicode code points contained in cutset removed.
//
// To remove a prefix, use TrimPrefix instead.
func TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// TrimRight returns a slice of the string s, with all trailing
// Unicode code points contained in cutset removed.
//
// To remove a suffix, use TrimSuffix instead.
func TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// TrimSpace returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimPrefixes returns the first of any of the prefixes that match returned
func TrimPrefixes(s string, prefixes []string) string {
	startLen := len(s)
	for _, pfx := range prefixes {
		s = TrimPrefix(s, pfx)
		if len(s) != startLen {
			return s
		}
	}
	return s
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// Replace returns a copy of the string s with the first n
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// ReplaceAll returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding.
func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// empty string if none
func First(strs []string, fns ...func(s string) bool) string {
	return StringSlice(strs).First(fns...)
}

func Filter(fn func(s string) bool, strs ...string) StringSlice {
	return StringSlice(strs).Filter(fn)
}

// Compact returns the first non-empty string argument in the list
func Compact(strs ...string) StringSlice {
	return StringSlice(strs).Compact()
}

// CleanAlphaNumeric returns a string with only A-Z, a-z, 0-9 and space allowed
func CleanAlphaNumeric(s string) string {
	if s == "" {
		return s
	}
	tmp := []byte(s)
	n := 0
	for _, b := range tmp {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			tmp[n] = b
			n++
		}
	}
	if n == 0 {
		return ""
	}
	return string(tmp[:n])

}
