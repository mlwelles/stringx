package stringx

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func NewStringSliceByQuotingItems(items ...interface{}) StringSlice {
	strs := make([]string, len(items))
	for i, item := range items {
		strs[i] = fmt.Sprintf("%q", item)
	}
	return strs
}

func (slice StringSlice) Less(i, j int) bool { return slice[i] < slice[j] }

// Sort is a convenience method.
func (slice StringSlice) Sort() { sort.Sort(slice) }

func (slice StringSlice) Sorted() StringSlice {
	s := slice.Copy()
	s.Sort()
	return s
}

func (slice StringSlice) SortedDesc() StringSlice {
	s := slice.Copy()
	s.Sort()
	return s
}

func (slice StringSlice) SortDesc() {
	sort.Sort(SortDescStringSlice(slice))
}

func (slice StringSlice) AnyEqual(s string) bool {
	return slice.Any(func(o string) bool {
		if s == o {
			return true
		}
		return false
	})
}

func (slice StringSlice) Join(sep string) string {
	return Join(slice, sep)
}

func (slice StringSlice) FindFuzzy(term string) StringSlice {
	return FindFuzzy(term, slice)
}

func (slice StringSlice) FindTermPrefix(term string) StringSlice {
	return FindTermPrefix(term, slice)
}

func (slice StringSlice) FindRegexp(re *regexp.Regexp) StringSlice {
	return FindRegexp(re, slice)
}

func (slice StringSlice) FindContainingSubstring(substr string) StringSlice {
	return FindContainingSubstring(substr, slice)
}

func (slice StringSlice) FindWith(finder Finder, term string) StringSlice {
	return FindWith(finder, term, slice)
}

func MapTrimSpace(slice StringSlice) StringSlice {
	return slice.Map(func(s string) string {
		return strings.TrimSpace(s)
	})
}

func (slice StringSlice) MapTrimSpace() StringSlice {
	return MapTrimSpace(slice)
}

type SortDescStringSlice []string

func (slice SortDescStringSlice) Len() int {
	return len(slice)
}

//actually greater!
func (slice SortDescStringSlice) Less(i, j int) bool {
	return slice[i] > slice[j]
}

func (slice SortDescStringSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
