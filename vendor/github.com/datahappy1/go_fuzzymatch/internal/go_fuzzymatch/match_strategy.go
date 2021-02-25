package gofuzzymatch

type strategy interface {
	matchStrings(s1 string, s2 string) int
}
