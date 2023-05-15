package common

// Misc utils

import (
	"net/http"
	"regexp"
)

var usernameRegex *regexp.Regexp = regexp.MustCompile(`^[0-9a-zA-Z_-]*[a-zA-Z0-9]+[0-9a-zA-Z_-]*$`)

const InvalidNameError string = `Invalid name.<br />Name must be between 2 and 36 characters in length; contain only numbers, letters, underscores or dashes; and contain at least one number or letter.<br />Names cannot contain spaces.`

// IsValidName checks that name is within the correct ranges, follows the regex defined
// and is not a valid color name
func IsValidName(name string) bool {
	return 2 <= len(name) && len(name) <= 36 &&
		usernameRegex.MatchString(name)
}

// Return the value of "Forwarded" or "X-Forwarded-For",
// if "Forwarded" & "X-Forwarded-For" are present then "Forwarded" value is returned.
// Return "" if "Forwarded" and "X-Forwarded-For" are absent.
func ExtractForwarded(r *http.Request) string {
	f := r.Header.Get("Forwarded")
	if f != "" {
		return f
	}

	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		return xff
	}

	return ""
}
