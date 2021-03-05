package repository

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
)

// Create returns string, bool
func Create(RequestID string, StringsToMatch []string, StringsToMatchIn []string,
	Mode string, RequestedFromIP string, BatchSize int) (string, bool) {
	fuzzyMatchObject := model.CreateFuzzyMatch(RequestID, StringsToMatch, StringsToMatchIn, Mode, RequestedFromIP, BatchSize, 0)
	model.RequestsData = append(model.RequestsData, fuzzyMatchObject)
	return "ok", true
}

// Update returns string, bool
func Update(requestID string, returnedRows int) (string, bool) {
	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {
			model.RequestsData[i] = model.UpdateFuzzyMatch(model.RequestsData[i], returnedRows)
			return "ok", true
		}
	}
	return "not found", false
}

// Delete returns string, bool
func Delete(requestID string) (string, bool) {
	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {
			model.RequestsData[i] = model.RequestsData[len(model.RequestsData)-1]
			model.RequestsData[len(model.RequestsData)-1] = model.FuzzyMatchModel{}
			model.RequestsData = model.RequestsData[:len(model.RequestsData)-1]
			return "ok", true
		}
	}
	return "not found", false
}

// GetByRequestID returns FuzzyMatchModel
func GetByRequestID(requestID string) model.FuzzyMatchModel {
	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {
			return model.CreateFuzzyMatch(
				model.RequestsData[i].RequestID,
				model.RequestsData[i].StringsToMatch,
				model.RequestsData[i].StringsToMatchIn,
				model.RequestsData[i].Mode,
				model.RequestsData[i].RequestedFromIP,
				model.RequestsData[i].BatchSize,
				model.RequestsData[i].ReturnedRows)
		}
	}
	return model.CreateDummyFuzzyMatch()
}

// GetByIP returns FuzzyMatchModel
func GetByIP(ip string) model.FuzzyMatchModel {
	for i := range model.RequestsData {
		if model.RequestsData[i].RequestedFromIP == ip {
			return model.CreateFuzzyMatch(
				model.RequestsData[i].RequestID,
				model.RequestsData[i].StringsToMatch,
				model.RequestsData[i].StringsToMatchIn,
				model.RequestsData[i].Mode,
				model.RequestsData[i].RequestedFromIP,
				model.RequestsData[i].BatchSize,
				model.RequestsData[i].ReturnedRows)
		}
	}
	return model.CreateDummyFuzzyMatch()
}

// CountAll returns int
func CountAll() int {
	return len(model.RequestsData)
}
