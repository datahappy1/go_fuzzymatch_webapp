package utils

import (
	"github.com/google/uuid"
	"strings"
)

// CreateUUID returns string
func CreateUUID() string {
	return uuid.New().String()
}

// IsValidUUID returns bool
func IsValidUUID(RequestUUID string) bool {
	_, err := uuid.Parse(RequestUUID)
	return err == nil
}

// IsStringInSlice returns bool
func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func safeCommaSplitter(str string, delimiter rune) []string {
	var isQuote = false

	f := func(c rune) bool {
		if c == '"' {
			if isQuote {
				isQuote = false
			} else {
				isQuote = true
			}
		}
		if !isQuote {
			return c == delimiter
		}
		return false
	}
	result := strings.FieldsFunc(str, f)
	return result
}

// SplitFormStringValueToSliceOfStrings returns array of strings
func SplitFormStringValueToSliceOfStrings(formValue string) []string {
	delimiter := []rune(",")[0]
	parsedStringsSlice := safeCommaSplitter(formValue, delimiter)
	return parsedStringsSlice
}
