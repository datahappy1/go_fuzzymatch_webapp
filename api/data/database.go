package data

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
)

var RequestsData = make(map[string]model.FuzzyMatchModel)

// InsertItem returns error
func InsertItem(fuzzyMatchObject model.FuzzyMatchModel) {
	RequestsData[fuzzyMatchObject.RequestID] = fuzzyMatchObject
}

// UpdateItem returns nil
func UpdateItem(requestID string, matchModel model.FuzzyMatchModel) {
	RequestsData[matchModel.RequestID] = matchModel
}

// DeleteItem returns nil
func DeleteItem(requestID string) {
	delete(RequestsData, requestID)
}

// GetItemByID returns FuzzyMatchModel
func GetItemByID(requestID string) model.FuzzyMatchModel {
	value, exists := RequestsData[requestID]

	if !exists {
		return model.FuzzyMatchModel{}
	}

	return value
}

// CountAllItems returns int
func CountAllItems() int {
	return len(RequestsData)
}

// GetAllItems returns map[string]model.FuzzyMatchModel
func GetAllItems() map[string]model.FuzzyMatchModel {
	return RequestsData
}
