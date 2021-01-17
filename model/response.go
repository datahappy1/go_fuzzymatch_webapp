package model

// FuzzyMatchResponse returns struct
type FuzzyMatchResponse struct {
	RequestID string
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
	NextToken string
}
