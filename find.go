//go:generate go run github.com/noho-digital/enumer -type=Finder

package stringx

import (
	lithammer "github.com/noho-digital/fuzzysearch/fuzzy"
	sahlim "github.com/sahilm/fuzzy"
	"regexp"
)

type Finder int

const (
	FinderUndefined Finder = iota
	FinderIndex
	FinderRegexp
	FinderTermPrefix
	FinderSahilmFuzzyFind
	// Find will return a list of strings in targets that fuzzy matches source.
	FinderLithammerFuzzyFind
	// FindFold is a case-insensitive version of Find.
	FinderLithammerFuzzyFindFold
	// FindNormalized is a unicode-normalized version of Find.
	FinderLithammerFuzzyFindNormalized
	// FindNormalizedFold is a unicode-normalized and case-insensitive version of Find.
	FinderLithammerFuzzyFindNormalizedFold
	// RankFind is similar to Find, except it will also rank all matches using
	// Levenshtein distance.
	FinderLithammerFuzzyRankFind
	// RankFindFold is a case-insensitive version of RankFind.
	FinderLithammerFuzzyRankFindFold
	// RankFindNormalized is a unicode-normalized version of RankFind.
	FinderLithammerFuzzyRankFindNormalized
	// RankFindNormalizedFold is a unicode-normalized and case-insensitive version of RankFind.
	FinderLithammerFuzzyRankFindNormalizedFold
)

var _startingDefaultFuzzyFinder = FinderLithammerFuzzyFindNormalizedFold
var _defaultFuzzyFinder = _startingDefaultFuzzyFinder
var _notAlphaNumericRe = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func (i Finder) IsFuzzy() bool {
	return Index(i.String(), "Fuzzy") > -1
}

func SetDefaultFuzzyFinder(t Finder) {
	if t.IsAFinder() && t.IsFuzzy() {
		_defaultFuzzyFinder = t
	}
}

func DefaultFuzzyFinder() Finder {
	if !_defaultFuzzyFinder.IsAFinder() || !_defaultFuzzyFinder.IsFuzzy() {
		_defaultFuzzyFinder = _startingDefaultFuzzyFinder
	}
	return _defaultFuzzyFinder
}

func Find(source string, targets StringSlice) StringSlice {
	return FindWith(DefaultFuzzyFinder(), source, targets)
}

func FindRegexp(re *regexp.Regexp, targets StringSlice) StringSlice {
	results := NewStringSlice()
	for _, target := range targets {
		if re != nil && re.MatchString(target) {
			results = append(results, target)
		}
	}
	return results
}

func FindFuzzy(source string, targets StringSlice) StringSlice {
	return FindWith(DefaultFuzzyFinder(), source, targets)
}

func FindContainingSubstring(source string, targets StringSlice) StringSlice {
	results := NewStringSlice()
	for _, target := range targets {
		if Index(target, source) >= 0 {
			results = append(results, target)
		}
	}
	return results
}

func FindTermPrefix(term string, targets StringSlice) StringSlice {
	results := NewStringSlice()
	ranks := lithammer.RankFindNormalizedFold(term, targets)
	cleanTerm := TrimSpace(ToLower(term))
	for _, rank := range ranks {
		match := false
		target := TrimSpace(ToLower(rank.Target))
		if Index(target, cleanTerm) == 0 {
			match = true
		} else {
			tokens := Fields(ToLower(rank.Target))
		Inner:
			for _, token := range tokens {
				for _, termToken := range Fields(cleanTerm) {
					if Index(token, termToken) == 0 {
						match = true
						break Inner
					}
					cleanToken := _notAlphaNumericRe.ReplaceAllString(token, "")
					if Index(cleanToken, termToken) == 0 {
						match = true
						break Inner
					}
				}
			}
		}
		if match {
			results = append(results, rank.Target)
		}
	}
	return results
}

func FindWith(finder Finder, source string, targets StringSlice) StringSlice {
	switch finder {
	case FinderUndefined, FinderIndex:
		return FindContainingSubstring(source, targets)
	case FinderRegexp:
		re, err := regexp.Compile(source)
		if err != nil {
			return NewStringSlice()
		}
		return FindRegexp(re, targets)
	case FinderTermPrefix:
		return FindTermPrefix(source, targets)
	case FinderSahilmFuzzyFind:
		return matchStrs(sahlim.Find(source, targets))
	case FinderLithammerFuzzyFind:
		return lithammer.Find(source, targets)
	case FinderLithammerFuzzyFindFold:
		return lithammer.FindFold(source, targets)
	case FinderLithammerFuzzyFindNormalized:
		return lithammer.FindNormalized(source, targets)
	case FinderLithammerFuzzyFindNormalizedFold:
		return lithammer.FindNormalizedFold(source, targets)
	case FinderLithammerFuzzyRankFind:
		return rankTargets(lithammer.RankFind(source, targets))
	case FinderLithammerFuzzyRankFindFold:
		return rankTargets(lithammer.RankFindFold(source, targets))
	case FinderLithammerFuzzyRankFindNormalized:
		return rankTargets(lithammer.RankFindNormalized(source, targets))
	case FinderLithammerFuzzyRankFindNormalizedFold:
		return rankTargets(lithammer.RankFindNormalizedFold(source, targets))
	default:
		return FindWith(_defaultFuzzyFinder, source, targets)
	}
}

func matchStrs(matches []sahlim.Match) StringSlice {
	ss := make(StringSlice, len(matches))
	for i, match := range matches {
		ss[i] = match.Str
	}
	return ss
}

func rankTargets(ranks lithammer.Ranks) StringSlice {
	ss := make(StringSlice, len(ranks))
	for i, rank := range ranks {
		ss[i] = rank.Target
	}
	return ss
}
