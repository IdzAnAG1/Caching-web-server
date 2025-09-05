package reg

import (
	"regexp"
)

const (
	QS = "QuantitySymbols"
	TL = "TwoLower"
	TU = "TwoUpper"
	OD = "OneDigit"
	SS = "SpecialSymbols"
)

var RegularExpression = map[string]string{
	QS: `^.{8,}$`,
	TL: `(?s).*(?:[a-z].*){2}`,
	TU: `(?s).*(?:[A-Z].*){2}`,
	OD: `\d`,
	SS: `[^a-zA-Z0-9]`,
}

func ValidatePassword(password string) bool {
	for _, value := range RegularExpression {
		if !Match(value, password) {
			return false
		}
	}
	return true
}

func ValidateLogin(login string) bool {
	for key, value := range RegularExpression {
		if key == QS {
			if !Match(value, login) {
				return false
			}
		}
		continue
	}
	return true
}

func Match(exp string, password string) bool {
	r := regexp.MustCompile(exp)
	return r.MatchString(password)
}
