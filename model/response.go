package model

// FuzzyMatchResponse returns struct
type FuzzyMatchResponse struct {
	RequestID string
}

// CreateFuzzyMatchResponse returns FuzzyMatchResponse
func CreateFuzzyMatchResponse(requestID string) FuzzyMatchResponse {
	resp := FuzzyMatchResponse{
		RequestID: requestID}
	return resp
}

// FuzzyMatchResult returns struct
type FuzzyMatchResult struct {
	StringToMatch string
	StringMatched string
	Result        uint16
}

// FuzzyMatchResults returns struct
type FuzzyMatchResults struct {
	Results []FuzzyMatchResult
}

// FuzzyMatchResultsResponse returns struct
type FuzzyMatchResultsResponse struct {
	RequestID string
	Mode      string
	Results   FuzzyMatchResults
}
