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

// AuxiliaryMatchResult returns struct
type AuxiliaryMatchResult struct {
	StringMatched string
	Result        int
}

// FuzzyMatchResult returns struct
type FuzzyMatchResult struct {
	StringToMatch string
	StringMatched string
	Result        int
}

// FuzzyMatchResultsResponse returns struct
type FuzzyMatchResultsResponse struct {
	RequestID string
	Mode      string
	Results   []FuzzyMatchResult
}
