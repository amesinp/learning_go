package utils

import (
	"strconv"
)

// ConvertToInt converts a string to an integer
// It retruns a fallback value if the string is not a valid integer
func ConvertToInt(value string, fallbackValue int) int {
	if value == "" {
		return fallbackValue
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return fallbackValue
	}

	return i
}
