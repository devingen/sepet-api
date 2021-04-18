package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"sort"
)

// isLetter checks if the given string is alphanumeric, contains - as the
// only non-alphanumeric char and is between 3 to 32 characters long.
var isLetter = regexp.MustCompile(`^[0-9a-zA-Z\-]{3,32}$`).MatchString

var reservedNames = []string{"api", "eu-central-1"}

// ValidateBucketDomain implements validator.Func
func ValidateBucketDomain(fl validator.FieldLevel) bool {
	domain := fl.Field().String()
	return isLetter(domain) && !contains(reservedNames, domain)
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
