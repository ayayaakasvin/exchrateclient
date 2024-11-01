package errorhand

import (
	"fmt"
)

// IfError checks if an error is not nil. If it is not nil, it returns a formatted error with a custom message.
// If the error is nil, it returns nil.
func IfError(msg string, err error) error {
	if err != nil {
		// If an error occurred, format and return a new error with the provided message and the original error message.
		return fmt.Errorf("%s: %s", msg, err.Error())
	}

	// If there was no error, return nil.
	return nil
} 