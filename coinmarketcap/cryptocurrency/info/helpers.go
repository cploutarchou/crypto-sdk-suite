package info

import (
	"fmt"
	"strings"
)

// SliceToString concatenates a slice of strings into a single string separated by the specified separator.
// Returns the concatenated string and an error if the array is empty.
func SliceToString(array []string, separator string) (string, error) {
	if len(array) == 0 {
		return "", fmt.Errorf("empty array")
	}

	finalVal := strings.Join(array, separator)
	return finalVal, nil
}

// BoolToString converts a boolean value to its string representation.
// Returns "true" if booleanValue is true, otherwise "false".
func BoolToString(booleanValue bool) string {
	if booleanValue {
		return "true"
	}
	return "false"
}
