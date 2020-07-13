package luhn

import (
	"unicode"
)

func Valid(number string) bool {
	num := []int{}
	for _, r := range number {
		if unicode.IsNumber(r) {
			num = append(num, int(r)-'0')
		} else if !unicode.IsSpace(r) {
			return false
		}
	}
	if len(num) < 1 || (len(num) == 1 && num[0] == 0) {
		return false
	}
	for i := len(num) - 2; i >= 0; i -= 2 {
		num[i] = 2 * num[i]
		if num[i] > 9 {
			num[i] -= 9
		}
	}
	sum := 0
	for _, i := range num {
		sum += i
	}
	return sum%10 == 0
}
