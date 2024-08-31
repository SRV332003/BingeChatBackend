package validators

import "regexp"

// IsValid method
func IsValidUsername(username string) bool {
	if len(username) < 5 || len(username) > 20 {
		return false
	}

	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return re.MatchString(username)

}
