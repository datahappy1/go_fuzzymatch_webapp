package controller

import (
	"strings"

	"github.com/google/uuid"
)

// FuzzyMatchExternalRequest returns struct
type FuzzyMatchExternalRequest struct {
	StringsToMatch   string `json:"stringsToMatch"`
	StringsToMatchIn string `json:"stringsToMatchIn"`
	Mode             string `json:"mode"`
}

// FuzzyMatchRequest returns struct
type FuzzyMatchRequest struct {
	RequestID        string
	StringsToMatch   []string
	StringsToMatchIn []string
	Mode             string
	RequestedFromIP  string
}

// CreateFuzzyMatchRequest returns FuzzyMatchRequest
func CreateFuzzyMatchRequest(stringsToMatch []string, stringsToMatchIn []string, mode string,
	requestedFromIP string) FuzzyMatchRequest {
	req := FuzzyMatchRequest{
		RequestID:        uuid.New().String(),
		StringsToMatch:   stringsToMatch,
		StringsToMatchIn: stringsToMatchIn,
		Mode:             mode,
		RequestedFromIP:  requestedFromIP}
	return req
}

// IsValidUUID returns UUID
func IsValidUUID(RequestUUID string) bool {
	_, err := uuid.Parse(RequestUUID)
	return err == nil
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
	return parsedStringsSlice
}
