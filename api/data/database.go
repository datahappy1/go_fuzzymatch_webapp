package data

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
)

type Database struct {
	RequestsData map[string]model.FuzzyMatchModel
}

// InsertItem returns error
func (d *Database) InsertItem(fuzzyMatchObject model.FuzzyMatchModel) {
	d.RequestsData[fuzzyMatchObject.RequestID] = fuzzyMatchObject
}

// UpdateItem returns nil
func (d *Database) UpdateItem(fuzzyMatchObject model.FuzzyMatchModel) {
	d.RequestsData[fuzzyMatchObject.RequestID] = fuzzyMatchObject
}

// DeleteItem returns nil
func (d *Database) DeleteItem(requestID string) {
	delete(d.RequestsData, requestID)
}

// GetItemByID returns FuzzyMatchModel
func (d *Database) GetItemByID(requestID string) model.FuzzyMatchModel {
	value, exists := d.RequestsData[requestID]

	if !exists {
		return model.FuzzyMatchModel{}
	}

	return value
}

// CountAllItems returns int
func (d *Database) CountAllItems() int {
	return len(d.RequestsData)
}

// GetAllItems returns map[string]model.FuzzyMatchModel
func (d *Database) GetAllItems() map[string]model.FuzzyMatchModel {
	return d.RequestsData
}
