package pkg

import (
	gofuzzymatch "github.com/datahappy1/go_fuzzymatch/internal/go_fuzzymatch"
)

// FuzzyMatch returns int
func FuzzyMatch(string1 string, string2 string, mode string) int {

	if string1 == string2 {
		return 100
	} else if string1 == "" || string2 == "" {
		return 0
	} else {
		var m = &gofuzzymatch.Match{}
		if mode == "simple" {
			m.Strategy = gofuzzymatch.Simple{}
			return m.MatchStrings(string1, string2)
		} else if mode == "deepDive" {
			m.Strategy = gofuzzymatch.DeepDive{}
			return m.MatchStrings(string1, string2)
		} else if mode == "combined" {
			m.Strategy = gofuzzymatch.Combined{}
			return m.MatchStrings(string1, string2)
		} else {
			return -1
		}
	}
}
