package model

import (
	"errors"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
)

// FuzzyMatchModel returns struct
type FuzzyMatchModel struct {
	RequestID              string
	StringsToMatch         []string
	StringsToMatchIn       []string
	Mode                   Mode
	RequestedOn            string
	BatchSize              int
	StringsToMatchLength   int
	StringsToMatchInLength int
	ReturnedRows           int
	ReturnedAllRows        bool
}

// CreateFuzzyMatch returns (FuzzyMatchModel, error)
func CreateFuzzyMatch(requestID string, stringsToMatch []string, stringsToMatchIn []string,
	mode string, batchSize int, returnedRows int) (FuzzyMatchModel, error) {

	modeEnumerated, ok := StringToMode(mode)
	if !ok {
		return FuzzyMatchModel{}, errors.New("cannot find mode")
	}

	return FuzzyMatchModel{
		RequestID:              requestID,
		StringsToMatch:         stringsToMatch,
		StringsToMatchIn:       stringsToMatchIn,
		Mode:                   modeEnumerated,
		RequestedOn:            utils.FormatTimestamp(utils.GetCurrentUTCTimestamp()),
		BatchSize:              batchSize,
		StringsToMatchLength:   len(stringsToMatch),
		StringsToMatchInLength: len(stringsToMatchIn),
		ReturnedRows:           returnedRows,
		ReturnedAllRows:        false}, nil
}
