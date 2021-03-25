package repository

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/data"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
)

// Create returns nil
func Create(db data.Database, RequestID string, StringsToMatch []string, StringsToMatchIn []string,
	Mode string, BatchSize int) {

	fuzzyMatchObject := model.CreateFuzzyMatch(
		RequestID,
		StringsToMatch,
		StringsToMatchIn,
		Mode,
		BatchSize,
		0)

	db.InsertItem(fuzzyMatchObject)
}

// Update returns nil
func Update(db data.Database, fuzzyMatchObject model.FuzzyMatchModel) {
	db.UpdateItem(fuzzyMatchObject)
}

// Delete returns nil
func Delete(db data.Database, requestID string) {
	db.DeleteItem(requestID)
}

// GetByRequestID returns FuzzyMatchModel
func GetByRequestID(db data.Database, requestID string) model.FuzzyMatchModel {
	item := db.GetItemByID(requestID)
	return item
}

// CountAll returns int
func CountAll(db data.Database) int {
	return db.CountAllItems()
}

// GetAllTimedOutRequestIDs returns []string
func GetAllTimedOutRequestIDs(db data.Database, RequestTTLInMinutes int) []string {
	var result []string
	currentDateTimeOffset := utils.GetCurrentDateTimeOffset(RequestTTLInMinutes)
	currentDateTimeOffsetUnixEpoch := utils.ConvertDateStringToUnixEpoch(currentDateTimeOffset)

	storedRequests := db.GetAllItems()

	for i := range storedRequests {
		if utils.ConvertDateStringToUnixEpoch(storedRequests[i].RequestedOn) <= currentDateTimeOffsetUnixEpoch {
			result = append(result, storedRequests[i].RequestID)
		}
	}
	return result
}
