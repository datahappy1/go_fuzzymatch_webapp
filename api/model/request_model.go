package model

import (
	"errors"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
	"strings"
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
}

// CreateFuzzyMatchRequest returns FuzzyMatchRequest
func CreateFuzzyMatchRequest(stringsToMatch []string, stringsToMatchIn []string, mode string) (*FuzzyMatchRequest, error) {

	if len(stringsToMatch) == 0 {
		return nil, errors.New("stringsToMatch is invalid")
	}

	if len(stringsToMatchIn) == 0 {
		return nil, errors.New("stringsToMatchIn is invalid")
	}

	capitalizedMode := strings.Title(strings.ToLower(mode))
	if _, ok := StringToModeMap[capitalizedMode]; !ok {
		return nil, errors.New("mode is invalid")
	}

	req := FuzzyMatchRequest{
		RequestID:        utils.CreateUUID(),
		StringsToMatch:   stringsToMatch,
		StringsToMatchIn: stringsToMatchIn,
		Mode:             mode}
	return &req, nil
}
