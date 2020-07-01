package raindrops

import (
	"strconv"
)

//definition of the function conv
func Convert(x int) (s string) {
	if x%3 != 0 && x%5 != 0 && x%7 != 0 {
		s = strconv.Itoa(x)
	} else {
		if x%3 == 0 {
			s = s + "Pling"
		}
		if x%5 == 0 {
			s = s + "Plang"
		}
		if x%7 == 0 {
			s = s + "Plong"
		}
	}
	return
}
