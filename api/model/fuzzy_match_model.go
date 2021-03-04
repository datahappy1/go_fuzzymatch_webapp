package model

import (
	"time"
)

// FuzzyMatchDAO returns struct
type FuzzyMatchDAO struct {
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

// CreateFuzzyMatchDAO returns FuzzyMatchDAO
func CreateFuzzyMatchDAO(requestID string, stringsToMatch []string, stringsToMatchIn []string,
	mode string, requestedFromIP string, batchSize int, returnedRows int) FuzzyMatchDAO {
	dao := FuzzyMatchDAO{
		RequestID:              requestID,
		StringsToMatch:         stringsToMatch,
		StringsToMatchIn:       stringsToMatchIn,
		Mode:                   mode,
		RequestedOn:            time.Now().String(),
		RequestedFromIP:        requestedFromIP,
		BatchSize:              batchSize,
		StringsToMatchLength:   len(stringsToMatch),
		StringsToMatchInLength: len(stringsToMatchIn),
		ReturnedRows:           returnedRows}
	return dao
}

// UpdateFuzzyMatchDAO returns FuzzyMatchDAO
func UpdateFuzzyMatchDAO(fuzzyMatchDAO FuzzyMatchDAO, returnedRows int) FuzzyMatchDAO {
	dao := fuzzyMatchDAO
	dao.ReturnedRows = returnedRows
	return dao
}
