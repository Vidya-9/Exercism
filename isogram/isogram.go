package isogram

import (
	"regexp"
	"strings"
)

//the definition of the function IsIsogram
func IsIsogram(str string) bool {
	if len(str) > 1 {
		str = strings.ToUpper(str)
		for _, v := range str {
			if (strings.Count(str, string(v)) > 1) && regexp.MustCompile("[A-Z]").MatchString(string(v)) {
				return false
			}
		}
	}
	return true
}
