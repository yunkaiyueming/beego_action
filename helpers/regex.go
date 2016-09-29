package helpers

import (
	"regexp"
)

func CheckEmail(str string) bool {
	pattern := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	return regrexMatch(str, pattern)
}

func CheckPhone(str string) bool {
	pattern := `^((\(\d{2,3}\))|(\d{3}\-))?13\d{9}$`
	return regrexMatch(str, pattern)
}

func CheckIp(str string) bool {
	pattern := `^(\d+)\.(\d+)\.(\d+)\.(\d+)$`
	return regrexMatch(str, pattern)
}

func CheckDate(str string) bool {
	pattern := `^(\d{4}|\d{2})-((1[0-2])|(0?[1-9]))-(([12][0-9])|(3[01])|(0?[1-9]))$`
	return regrexMatch(str, pattern)
}

func CheckUrl(str string) bool {
	pattern := `^[a-zA-z]+://[^\s]*$`
	return regrexMatch(str, pattern)
}

func CheckChinese(str string) bool {
	pattern := `^[\u0391-\uFFE5]+$`
	return regrexMatch(str, pattern)
}

func CheckNum(str string) bool {
	pattern := `^-?\d+$`
	return regrexMatch(str, pattern)
}

func regrexMatch(str, pattern string) bool {
	ret, _ := regexp.MatchString(pattern, str)
	return ret
}
