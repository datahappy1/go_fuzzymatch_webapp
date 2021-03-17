package repository

import (
	"errors"
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

	model.RequestsData = append(model.RequestsData, fuzzyMatchObject)
	return nil
}

// Update returns (string, error)
func Update(requestID string, returnedRows int) error {
	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {
			model.RequestsData[i] = model.UpdateFuzzyMatch(model.RequestsData[i], returnedRows)
			return nil
		}
	}
	return errors.New("request not found, not updated")
}

// Delete returns error
func Delete(requestID string) error {
	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {
			model.RequestsData[i] = model.RequestsData[len(model.RequestsData)-1]
			model.RequestsData[len(model.RequestsData)-1] = model.FuzzyMatchModel{}
			model.RequestsData = model.RequestsData[:len(model.RequestsData)-1]
			return nil
		}
	}
	return errors.New("request not found, not deleted")
}

// GetByRequestID returns FuzzyMatchModel
func GetByRequestID(requestID string) model.FuzzyMatchModel {
	for i := range model.RequestsData {
		if model.RequestsData[i].RequestID == requestID {
			return model.RequestsData[i]
		}
	}
	return model.FuzzyMatchModel{}
}

// GetAll returns []model.FuzzyMatchModel
func GetAll() []model.FuzzyMatchModel {
	return model.RequestsData
}

// CountAll returns int
func CountAll() int {
	return len(model.RequestsData)
}

// GetAllTimedOutRequests returns []model.FuzzyMatchModel
func GetAllTimedOutRequests(RequestTTLInMinutes int) []model.FuzzyMatchModel {
	var result []model.FuzzyMatchModel
	currentDateTimeOffset := utils.GetCurrentDateTimeOffset(RequestTTLInMinutes)
	currentDateTimeOffsetUnixEpoch := utils.ConvertDateStringToUnixEpoch(currentDateTimeOffset)

	for i := range model.RequestsData {
		if utils.ConvertDateStringToUnixEpoch(model.RequestsData[i].RequestedOn) <= currentDateTimeOffsetUnixEpoch {
			result = append(result, model.RequestsData[i])
		}
	}
	return result
}
