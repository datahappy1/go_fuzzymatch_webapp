package model

import (
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
func CreateFuzzyMatchRequest(stringsToMatch []string, stringsToMatchIn []string, mode string) *FuzzyMatchRequest {
	req := FuzzyMatchRequest{
		RequestID:        uuid.New().String(),
		StringsToMatch:   stringsToMatch,
		StringsToMatchIn: stringsToMatchIn,
		Mode:             mode}
	return &req
}
