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
	flag := true
	for _, value := range RegularExpression {
		flag = Match(value, password)
	}
	return flag
}

func ValidateLogin(login string) bool {
	flag := true
	for key, value := range RegularExpression {
		if key == QS {
			flag = Match(value, login)
		}
		continue
	}
	return flag
}

func Match(exp string, password string) bool {
	r := regexp.MustCompile(exp)
	return r.MatchString(password)
}
