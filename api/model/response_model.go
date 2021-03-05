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

// FuzzyMatchResults returns []FuzzyMatchResult
type FuzzyMatchResults []FuzzyMatchResult

// FuzzyMatchResultsResponse returns struct
type FuzzyMatchResultsResponse struct {
	RequestID       string
	Mode            string
	RequestedOn     string
	ReturnedAllRows bool
	Results         FuzzyMatchResults
}

// CreateFuzzyMatchResultsResponse returns FuzzyMatchResultsResponse
func CreateFuzzyMatchResultsResponse(
	requestID, mode, requestedOn string,
	returnedAllRows bool,
	results FuzzyMatchResults) FuzzyMatchResultsResponse {
	return FuzzyMatchResultsResponse{
		RequestID:       requestID,
		Mode:            mode,
		RequestedOn:     requestedOn,
		ReturnedAllRows: returnedAllRows,
		Results:         results,
	}
}
