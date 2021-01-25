package model

import "time"

// FuzzyMatchDAO returns struct
type FuzzyMatchDAO struct {
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

// CreateFuzzyMatchDAO returns FuzzyMatchDAO
func CreateFuzzyMatchDAO(requestID string, stringsToMatch []string, stringsToMatchIn []string, mode string) FuzzyMatchDAO {
	dao := FuzzyMatchDAO{
		RequestID:              requestID,
		StringsToMatch:         stringsToMatch,
		StringsToMatchIn:       stringsToMatchIn,
		Mode:                   mode,
		RequestedOn:            time.Now().String(),
		BatchSize:              2,
		StringsToMatchLength:   len(stringsToMatch),
		StringsToMatchInLength: len(stringsToMatchIn),
		ReturnedRows:           0}
	return dao
}

// CreateFuzzyMatchDAOInRequestsData returns FuzzyMatchDAO
func CreateFuzzyMatchDAOInRequestsData(RequestID string, StringsToMatch []string, StringsToMatchIn []string, Mode string) (string, bool) {
	dao := CreateFuzzyMatchDAO(RequestID, StringsToMatch, StringsToMatchIn, Mode)
	RequestsData = append(RequestsData, dao)
	return "ok", true
}

// UpdateFuzzyMatchDAO returns FuzzyMatchDAO
func UpdateFuzzyMatchDAO(fuzzyMatchDAO FuzzyMatchDAO, returnedRows int) FuzzyMatchDAO {
	dao := fuzzyMatchDAO
	dao.ReturnedRows = returnedRows
	return dao
}

// UpdateFuzzyMatchDAOInRequestsData returns FuzzyMatchDAO
func UpdateFuzzyMatchDAOInRequestsData(requestID string, returnedRows int) (string, bool) {
	for i := range RequestsData {
		if RequestsData[i].RequestID == requestID {
			RequestsData[i] = UpdateFuzzyMatchDAO(RequestsData[i], returnedRows)
			break
		}
	}
	return "ok", true
}

// DeleteFuzzyMatchDAOInRequestsData returns FuzzyMatchDAO
func DeleteFuzzyMatchDAOInRequestsData(requestID string) (string, bool) {
	for i := range RequestsData {
		if RequestsData[i].RequestID == requestID {
			RequestsData[i] = RequestsData[len(RequestsData)-1]
			RequestsData[len(RequestsData)-1] = FuzzyMatchDAO{}
			RequestsData = RequestsData[:len(RequestsData)-1]
			break
		}
	}
	return "ok", true
}

// RequestsData returns []FuzzyMatchDAO
var RequestsData []FuzzyMatchDAO
