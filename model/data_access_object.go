package model

import "time"

// FuzzyMatchDAO returns struct
type FuzzyMatchDAO struct {
	RequestID            string
	StringsToMatch       []string
	StringsToMatchIn     []string
	Mode                 string
	RequestedOn          string
	BatchSize            int
	StringsToMatchLength int
	ReturnedRows         int
	ReturnedAllRows      bool
}

// CreateFuzzyMatchDAO returns FuzzyMatchDAO
func CreateFuzzyMatchDAO(requestID string, stringsToMatch []string, stringsToMatchIn []string, mode string) FuzzyMatchDAO {
	dao := FuzzyMatchDAO{
		RequestID:            requestID,
		StringsToMatch:       stringsToMatch,
		StringsToMatchIn:     stringsToMatchIn,
		Mode:                 mode,
		RequestedOn:          time.Now().String(),
		BatchSize:            100,
		StringsToMatchLength: len(stringsToMatch),
		ReturnedRows:         0,
		ReturnedAllRows:      false}
	return dao
}
