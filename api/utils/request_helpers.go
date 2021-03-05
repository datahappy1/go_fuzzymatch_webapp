package utils

import (
	"fmt"
	"github.com/google/uuid"
	"net"
	"net/http"
	"strings"
)

// GetIP returns (string, error)
func GetIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("no valid ip found")
}

// CreateUUID returns string
func CreateUUID() string {
	return uuid.New().String()
}

// IsValidUUID returns bool
func IsValidUUID(RequestUUID string) bool {
	_, err := uuid.Parse(RequestUUID)
	return err == nil
}

// IsStringInSlice returns bool
func IsStringInSlice(a string, list [3]string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func safeCommaSplitter(str string, delimiter rune) []string {
	var isQuote = false

	f := func(c rune) bool {
		if c == '\'' {
			if isQuote {
				isQuote = false
			} else {
				isQuote = true
			}
		}
		if !isQuote {
			return c == delimiter
		}
		return false
	}
	result := strings.FieldsFunc(str, f)
	return result
}

// SplitFormStringValueToSliceOfStrings returns array of strings
func SplitFormStringValueToSliceOfStrings(formValue string) []string {
	delimiter := []rune(",")[0]
	parsedStringsSlice := safeCommaSplitter(formValue, delimiter)
	return parsedStringsSlice
}
