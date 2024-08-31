package validators

import (
	"errors"
	"regexp"
	"strings"
)

func IsValidEmail(email string) error {
	if len(email) >= 50 {
		return errors.New("too long email address")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	hasSpace := regexp.MustCompile(`\s+`).MatchString(email)
	if !re.MatchString(email) || hasSpace {
		return errors.New("bad email format")
	}

	return nil

}

func NormalizeEmail(email string) string {
	slug := strings.Split(email, "@")[0]
	slug = strings.Split(slug, "+")[0]
	email = slug + "@" + strings.Split(email, "@")[1]
	return strings.ToLower(email)
}
