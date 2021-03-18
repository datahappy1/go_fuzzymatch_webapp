package repository

import (
	"errors"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/data"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/utils"
)

// Create returns string
func Create(RequestID string, StringsToMatch []string, StringsToMatchIn []string,
	Mode string, BatchSize int) error {

	fuzzyMatchObject := model.CreateFuzzyMatch(
		RequestID,
		StringsToMatch,
		StringsToMatchIn,
		Mode,
		BatchSize,
		0)

	err := data.InsertItem(fuzzyMatchObject)
	if err != nil {
		return errors.New("request not inserted to database")
	}
	return nil
}

// Update returns error
func Update(requestID string, matchModel model.FuzzyMatchModel) error {
	err := data.UpdateItem(requestID, matchModel)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Delete returns error
func Delete(requestID string) error {
	err := data.DeleteItem(requestID)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
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
