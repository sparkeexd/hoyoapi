package utilities

import (
	"fmt"
	"os"

	"github.com/sparkeexd/hoyoapi/internal/errors"
)

// Parsed error structure.
type parseError struct {
	message string
}

// Returns the error message during environment variable parsing.
func (e parseError) Error() string {
	return e.message
}

// Helper function for reading strings from the environment.
func ParseString(input string) (string, error) {
	return input, nil
}

// Gets the parsed environment variable.
// Panics if the value is not found.
func GetEnv[V any](envKey string, parse func(string) (V, error)) V {
	value, err := getEnv(envKey, parse)
	if err != nil {
		panic(err)
	}

	return value
}

// Attempts to return the value of the specified env variable, otherwise return an error.
func getEnv[V any](envKey string, parse func(string) (V, error)) (V, error) {
	rawValue, found := os.LookupEnv(envKey)
	if !found {
		var noResult V
		return noResult, errors.NewError(
			errors.ENV_VARIABLE_READ_ERROR,
			fmt.Sprintf("Environment variable not set: %s", envKey),
		)
	}

	parsedValue, err := parse(rawValue)
	if err != nil {
		var noResult V
		return noResult, errors.NewError(
			errors.ENV_VARIABLE_PARSE_ERROR,
			fmt.Sprintf("Parsing %s value: %v", envKey, err),
		)
	}

	return parsedValue, nil
}
