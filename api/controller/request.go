package controller

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var modeTypes = [3]string{"simple", "deepDive", "combined"}

func stringInSlice(a string, list [3]string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

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
	requestedFromIP string) (*FuzzyMatchRequest, error) {

	if len(stringsToMatch) == 0 {
		return nil, errors.New("stringsToMatch is invalid")
	}

	if len(stringsToMatchIn) == 0 {
		return nil, errors.New("stringsToMatchIn is invalid")
	}

	if mode == "" || stringInSlice(mode, modeTypes) == false {
		return nil, errors.New("mode is invalid")
	}

	req := FuzzyMatchRequest{
		RequestID:        uuid.New().String(),
		StringsToMatch:   stringsToMatch,
		StringsToMatchIn: stringsToMatchIn,
		Mode:             mode,
		RequestedFromIP:  requestedFromIP}
	return &req, nil
}

// IsValidUUID returns bool
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
