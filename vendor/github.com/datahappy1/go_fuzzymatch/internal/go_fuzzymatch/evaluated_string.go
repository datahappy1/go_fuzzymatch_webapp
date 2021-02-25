package gofuzzymatch

import (
	"log"
	"regexp"
	"strings"
)

type evaluatedString struct {
	value                       string
	valueByWordSplitSlice       []string
	valueByWordSplitSliceLength int
}

func createEvaluatedString(v string) *evaluatedString {
	processedInputString := prepareString(v)
	stringWordSplit := splitStringToUniqueValuesSliceByWhitespace(processedInputString)
	s := evaluatedString{
		value:                       processedInputString,
		valueByWordSplitSlice:       stringWordSplit,
		valueByWordSplitSliceLength: len(stringWordSplit)}
	return &s
}

func toLower(s string) string {
	return strings.ToLower(s)
}

func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func removeUnusedChars(s string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(s, " ")
}

func prepareString(s string) string {
	return toLower(trimSpace(removeUnusedChars(s)))
}

func appendStringToSliceIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func splitStringToUniqueValuesSliceByWhitespace(s string) []string {
	var splitStringSlice []string
	for _, word := range strings.Fields(s) {
		splitStringSlice = appendStringToSliceIfMissing(splitStringSlice, word)
	}
	return splitStringSlice
}
