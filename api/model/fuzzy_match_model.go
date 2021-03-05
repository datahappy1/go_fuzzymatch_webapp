package model

import (
	"time"
)

// FuzzyMatchModel returns struct
type FuzzyMatchModel struct {
	RequestID              string
	StringsToMatch         []string
	StringsToMatchIn       []string
	Mode                   string
	RequestedOn            string
	RequestedFromIP        string
	BatchSize              int
	StringsToMatchLength   int
	StringsToMatchInLength int
	ReturnedRows           int
	ReturnedAllRows        bool
}

// CreateDummyFuzzyMatch returns FuzzyMatchModel
func CreateDummyFuzzyMatch() FuzzyMatchModel {
	return FuzzyMatchModel{
		RequestID:              "",
		StringsToMatch:         []string{},
		StringsToMatchIn:       []string{},
		Mode:                   "",
		RequestedOn:            "",
		RequestedFromIP:        "",
		BatchSize:              0,
		StringsToMatchLength:   0,
		StringsToMatchInLength: 0,
		ReturnedRows:           0,
		ReturnedAllRows:        false}
}

// CreateFuzzyMatch returns FuzzyMatchModel
func CreateFuzzyMatch(requestID string, stringsToMatch []string, stringsToMatchIn []string,
	mode string, requestedFromIP string, batchSize int, returnedRows int) FuzzyMatchModel {
	return FuzzyMatchModel{
		RequestID:              requestID,
		StringsToMatch:         stringsToMatch,
		StringsToMatchIn:       stringsToMatchIn,
		Mode:                   mode,
		RequestedOn:            time.Now().String(),
		RequestedFromIP:        requestedFromIP,
		BatchSize:              batchSize,
		StringsToMatchLength:   len(stringsToMatch),
		StringsToMatchInLength: len(stringsToMatchIn),
		ReturnedRows:           returnedRows,
		ReturnedAllRows:        false}
}

// UpdateFuzzyMatch returns FuzzyMatchModel
func UpdateFuzzyMatch(fuzzyMatch FuzzyMatchModel, returnedRows int) FuzzyMatchModel {
	updatedFuzzyMatch := fuzzyMatch
	updatedFuzzyMatch.ReturnedRows = returnedRows
	return updatedFuzzyMatch
}

// RequestsData returns []FuzzyMatchModel
var RequestsData []FuzzyMatchModel
