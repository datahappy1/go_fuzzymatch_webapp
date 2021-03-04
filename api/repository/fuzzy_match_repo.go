package repository

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
)

// CreateFuzzyMatchDAOInRequestsData returns FuzzyMatchDAO
func CreateFuzzyMatchDAOInRequestsData(RequestID string, StringsToMatch []string, StringsToMatchIn []string,
	Mode string, RequestedFromIP string, BatchSize int) (string, bool) {
	dao := model.CreateFuzzyMatchDAO(RequestID, StringsToMatch, StringsToMatchIn, Mode, RequestedFromIP, BatchSize, 0)
	RequestsData = append(RequestsData, dao)
	return "ok", true
}

// UpdateFuzzyMatchDAOInRequestsData returns FuzzyMatchDAO
func UpdateFuzzyMatchDAOInRequestsData(requestID string, returnedRows int) (string, bool) {
	for i := range RequestsData {
		if RequestsData[i].RequestID == requestID {
			RequestsData[i] = model.UpdateFuzzyMatchDAO(RequestsData[i], returnedRows)
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
			RequestsData[len(RequestsData)-1] = model.FuzzyMatchDAO{}
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
func EvaluateRequestCount(activeRequestsCount int) bool {
	if len(RequestsData) >= activeRequestsCount {
		return false
	}
	return true
}

// RequestsData returns []FuzzyMatchDAO
var RequestsData []model.FuzzyMatchDAO
