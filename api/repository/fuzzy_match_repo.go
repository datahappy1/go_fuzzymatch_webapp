package repository

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/data"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
)

// Create returns nil
func Create(RequestID string, StringsToMatch []string, StringsToMatchIn []string,
	Mode string, BatchSize int) {

	fuzzyMatchObject := model.CreateFuzzyMatch(
		RequestID,
		StringsToMatch,
		StringsToMatchIn,
		Mode,
		BatchSize,
		0)

	data.InsertItem(fuzzyMatchObject)
}

// Update returns nil
func Update(requestID string, matchModel model.FuzzyMatchModel) {
	data.UpdateItem(requestID, matchModel)
}

// Delete returns nil
func Delete(requestID string) {
	data.DeleteItem(requestID)
}

// GetByRequestID returns FuzzyMatchModel
func GetByRequestID(requestID string) model.FuzzyMatchModel {
	item := data.GetItemByID(requestID)
	return item
}

// CountAll returns int
func CountAll() int {
	return data.CountAllItems()
}

// GetAllTimedOutRequestIDs returns []string
func GetAllTimedOutRequestIDs(RequestTTLInMinutes int) []string {
	var result []string
	currentDateTimeOffset := utils.GetCurrentDateTimeOffset(RequestTTLInMinutes)
	currentDateTimeOffsetUnixEpoch := utils.ConvertDateStringToUnixEpoch(currentDateTimeOffset)

	storedRequests := data.GetAllItems()

	for i := range storedRequests {
		if utils.ConvertDateStringToUnixEpoch(storedRequests[i].RequestedOn) <= currentDateTimeOffsetUnixEpoch {
			result = append(result, storedRequests[i].RequestID)
		}
	}
	return result
}
