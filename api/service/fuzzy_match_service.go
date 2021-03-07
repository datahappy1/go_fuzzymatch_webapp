package service

import (
	fm "github.com/datahappy1/go_fuzzymatch/pkg"
	"github.com/datahappy1/go_fuzzymatch_webapp/api/model"
	"sort"
)

// CalculateFuzzyMatchingResults returns model.FuzzyMatchResults, bool, int
func CalculateFuzzyMatchingResults(fuzzyMatchObject model.FuzzyMatchModel) (model.FuzzyMatchResults, bool, int) {

	var fuzzyMatchResults model.FuzzyMatchResults
	var returnedAllRows bool
	var returnedRowsUpperBound int

	if fuzzyMatchObject.ReturnedRows + fuzzyMatchObject.BatchSize >= fuzzyMatchObject.StringsToMatchLength {
		returnedRowsUpperBound = fuzzyMatchObject.StringsToMatchLength
		returnedAllRows = true
	} else {
		returnedRowsUpperBound = fuzzyMatchObject.ReturnedRows + fuzzyMatchObject.BatchSize
		returnedAllRows = false
	}

	for stringToMatch := fuzzyMatchObject.ReturnedRows; stringToMatch < returnedRowsUpperBound; stringToMatch++ {
		var auxiliaryMatchResults []model.AuxiliaryMatchResult

		for stringToMatchIn := 0; stringToMatchIn < fuzzyMatchObject.StringsToMatchInLength; stringToMatchIn++ {
			auxiliaryMatchResult := model.AuxiliaryMatchResult{
				StringMatched: fuzzyMatchObject.StringsToMatchIn[stringToMatchIn],
				Result: fm.FuzzyMatch(
					fuzzyMatchObject.StringsToMatch[stringToMatch],
					fuzzyMatchObject.StringsToMatchIn[stringToMatchIn],
					fuzzyMatchObject.Mode)}

			auxiliaryMatchResults = append(auxiliaryMatchResults, auxiliaryMatchResult)
		}

		sort.SliceStable(auxiliaryMatchResults, func(i, j int) bool {
			return auxiliaryMatchResults[i].Result > auxiliaryMatchResults[j].Result
		})

		fuzzyMatchResult := model.FuzzyMatchResult{
			StringToMatch: fuzzyMatchObject.StringsToMatch[stringToMatch],
			StringMatched: auxiliaryMatchResults[0].StringMatched,
			Result:        auxiliaryMatchResults[0].Result}

		fuzzyMatchResults = append(fuzzyMatchResults, fuzzyMatchResult)
	}
	return fuzzyMatchResults, returnedAllRows, returnedRowsUpperBound
}
