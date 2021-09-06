package model

import "strings"

// Mode is int
type Mode int

const (
	Simple Mode = iota
	Deepdive
	Combined
)

// String returns string
func (s Mode) String() string {
	return [...]string{strings.ToLower("Simple"), strings.ToLower("Deepdive"), strings.ToLower("Combined")}[s]
}

var StringToModeMap = map[string]Mode{
	"Simple":   Simple,
	"Deepdive": Deepdive,
	"Combined": Combined,
}

// StringToMode returns (Mode, bool)
func StringToMode(s string) (Mode, bool) {
	val, ok := StringToModeMap[strings.Title(strings.ToLower(s))]
	return val, ok
}
