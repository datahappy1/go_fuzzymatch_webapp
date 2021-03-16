package model

import (
	"errors"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
)

var modeTypes = []string{"simple", "deepDive", "combined"}

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

	if mode == "" || utils.IsStringInSlice(mode, modeTypes) == false {
		return nil, errors.New("mode is invalid")
	}

	req := FuzzyMatchRequest{
		RequestID:        utils.CreateUUID(),
		StringsToMatch:   stringsToMatch,
		StringsToMatchIn: stringsToMatchIn,
		Mode:             mode,
		RequestedFromIP:  requestedFromIP}
	return &req, nil
}
