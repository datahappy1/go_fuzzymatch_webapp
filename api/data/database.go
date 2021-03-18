package data

import (
	"errors"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
)

// RequestsData returns []FuzzyMatchModel
var RequestsData []model.FuzzyMatchModel

// InsertItem returns error
func InsertItem(fuzzyMatchObject model.FuzzyMatchModel) error {
	RequestsData = append(RequestsData, fuzzyMatchObject)
	return nil
}

// UpdateItem returns error
func UpdateItem(requestID string, matchModel model.FuzzyMatchModel) error {
	for i := range RequestsData {
		if RequestsData[i].RequestID == requestID {
			RequestsData[i] = matchModel
			return nil
		}
	}
	return errors.New("request not found, not updated")
}

// DeleteItem returns error
func DeleteItem(requestID string) error {
	for i := range RequestsData {
		if RequestsData[i].RequestID == requestID {
			RequestsData[i] = RequestsData[len(RequestsData)-1]
			RequestsData[len(RequestsData)-1] = model.FuzzyMatchModel{}
			RequestsData = RequestsData[:len(RequestsData)-1]
			return nil
		}
	}
	return errors.New("request not found, not deleted")
}

// GetItemByID returns FuzzyMatchModel
func GetItemByID(requestID string) model.FuzzyMatchModel {
	for i := range RequestsData {
		if RequestsData[i].RequestID == requestID {
			return RequestsData[i]
		}
	}
	return model.FuzzyMatchModel{}
}

// CountAllItems returns int
func CountAllItems() int {
	return len(RequestsData)
}

// GetAllItems returns RequestsData
func GetAllItems() []model.FuzzyMatchModel {
	return RequestsData
}
