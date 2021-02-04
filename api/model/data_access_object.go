package model

import (
	"time"
)

var batchSize = 20
var activeRequestsCount = 1000

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
	mode string, requestedFromIP string, returnedRows int) FuzzyMatchDAO {
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

// CreateFuzzyMatchDAOInRequestsData returns FuzzyMatchDAO
func CreateFuzzyMatchDAOInRequestsData(RequestID string, StringsToMatch []string, StringsToMatchIn []string,
	Mode string, RequestedFromIP string) (string, bool) {
	dao := CreateFuzzyMatchDAO(RequestID, StringsToMatch, StringsToMatchIn, Mode, RequestedFromIP, 0)
	RequestsData = append(RequestsData, dao)
	return "ok", true
}

func updateFuzzyMatchDAO(fuzzyMatchDAO FuzzyMatchDAO, returnedRows int) FuzzyMatchDAO {
	dao := fuzzyMatchDAO
	dao.ReturnedRows = returnedRows
	return dao
}

// UpdateFuzzyMatchDAOInRequestsData returns FuzzyMatchDAO
func UpdateFuzzyMatchDAOInRequestsData(requestID string, returnedRows int) (string, bool) {
	for i := range RequestsData {
		if RequestsData[i].RequestID == requestID {
			RequestsData[i] = updateFuzzyMatchDAO(RequestsData[i], returnedRows)
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

// EvaluateRequestRatePerIP returns bool
func EvaluateRequestRatePerIP(ip string) (bool, string) {
	for i := range RequestsData {
		if RequestsData[i].RequestedFromIP == ip {
			return true, RequestsData[i].RequestID
		}
	}
	return false, ""
}

// EvaluateRequestCount returns bool
func EvaluateRequestCount() bool {
	if len(RequestsData) > activeRequestsCount {
		return false
	}
	return true
}

// RequestsData returns []FuzzyMatchDAO
var RequestsData []FuzzyMatchDAO
