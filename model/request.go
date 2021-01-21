package model

import (
	"strings"

	"github.com/google/uuid"
)

// FuzzyMatchRequest returns struct
type FuzzyMatchRequest struct {
	RequestID        string
	StringsToMatch   []string
	StringsToMatchIn []string
	Mode             string
}

// CreateFuzzyMatchRequest returns FuzzyMatchRequest
func CreateFuzzyMatchRequest(stringsToMatch []string, stringsToMatchIn []string, mode string) FuzzyMatchRequest {
	req := FuzzyMatchRequest{
		RequestID:        uuid.New().String(),
		StringsToMatch:   stringsToMatch,
		StringsToMatchIn: stringsToMatchIn,
		Mode:             mode}
	return req
}

func safeCommaSplitter(str string, delimiter rune) []string {
	var isQuote = false

	f := func(c rune) bool {
		if c == '\'' {
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
	//fmt.Println(len(formValue), formValue, o, len(o), w)

	//return strings.Split(formValue, ",")
	return parsedStringsSlice
}
