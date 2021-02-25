package gofuzzymatch

// Match type
type Match struct {
	Strategy strategy
}

// MatchStrings returns int
func (m *Match) MatchStrings(s1 string, s2 string) int {
	return m.Strategy.matchStrings(s1, s2)
}
