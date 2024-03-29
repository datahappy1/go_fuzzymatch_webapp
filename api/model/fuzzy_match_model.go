package model

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
)

// FuzzyMatchModel returns struct
type FuzzyMatchModel struct {
	RequestID              string
	StringsToMatch         []string
	StringsToMatchIn       []string
	Mode                   string
	RequestedOn            string
	BatchSize              int
	StringsToMatchLength   int
	StringsToMatchInLength int
	ReturnedRows           int
	ReturnedAllRows        bool
}

// CreateFuzzyMatch returns FuzzyMatchModel
func CreateFuzzyMatch(requestID string, stringsToMatch []string, stringsToMatchIn []string,
	mode string, batchSize int, returnedRows int) FuzzyMatchModel {

	return FuzzyMatchModel{
		RequestID:              requestID,
		StringsToMatch:         stringsToMatch,
		StringsToMatchIn:       stringsToMatchIn,
		Mode:                   mode,
		RequestedOn:            utils.FormatTimestamp(utils.GetCurrentUTCTimestamp()),
		BatchSize:              batchSize,
		StringsToMatchLength:   len(stringsToMatch),
		StringsToMatchInLength: len(stringsToMatchIn),
		ReturnedRows:           returnedRows,
		ReturnedAllRows:        false}
}
