package data

import (
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
)

// Database is struct
type Database struct {
	RequestsPseudoTable map[string]model.FuzzyMatchModel
}

// CreateRequestsPseudoTable returns map[string]model.FuzzyMatchModel
func CreateRequestsPseudoTable() map[string]model.FuzzyMatchModel {
	return make(map[string]model.FuzzyMatchModel)
}

// InsertItem returns nil
func (d *Database) InsertItem(fuzzyMatchObject model.FuzzyMatchModel) {
	d.RequestsPseudoTable[fuzzyMatchObject.RequestID] = fuzzyMatchObject
}

// UpdateItem returns nil
func (d *Database) UpdateItem(fuzzyMatchObject model.FuzzyMatchModel) {
	d.RequestsPseudoTable[fuzzyMatchObject.RequestID] = fuzzyMatchObject
}

// DeleteItem returns nil
func (d *Database) DeleteItem(requestID string) {
	delete(d.RequestsPseudoTable, requestID)
}

// GetItemByID returns FuzzyMatchModel
func (d *Database) GetItemByID(requestID string) model.FuzzyMatchModel {
	value, exists := d.RequestsPseudoTable[requestID]

	if !exists {
		return model.FuzzyMatchModel{}
	}

	return value
}

// CountAllItems returns int
func (d *Database) CountAllItems() int {
	return len(d.RequestsPseudoTable)
}

// GetAllItems returns map[string]model.FuzzyMatchModel
func (d *Database) GetAllItems() map[string]model.FuzzyMatchModel {
	return d.RequestsPseudoTable
}
