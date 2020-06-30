package hamming

import (
	"errors"
)

//the definition of the function "Distance()..."
func Distance(a, b string) (int, error) {
	var x int = 0
	var y error = nil

	if len(a) == len(b) {

		for i := range a {
			if a[i] != b[i] {
				x++
			}
		}
		return x, y
	}
	y = errors.New("String must have equal lengths")
	return -1, y
}
