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
		BatchSize:            2,
		StringsToMatchLength: len(stringsToMatch),
		ReturnedRows:         0,
		ReturnedAllRows:      false}
	return dao
}

// UpdateFuzzyMatchDAO returns FuzzyMatchDAO
func UpdateFuzzyMatchDAO(requests []FuzzyMatchDAO, requestID string, returnedRows int, returnedAllRows bool) FuzzyMatchDAO {
	var dao FuzzyMatchDAO
	for i := range requests {
		if requests[i].RequestID == requestID {
			dao = requests[i]

			dao.ReturnedRows = returnedRows
			dao.ReturnedAllRows = returnedAllRows

			requests[i] = dao

			break
		}
	}

	return dao
}
