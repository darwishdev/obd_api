package util

import "strings"

// checks if a given string str is present in an array of strings arr.
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if strings.Contains(a, str) {
			return true
		}
	}
	return false
}
