package scrabble

import (
	"strings"
)

//the body of the function Scrabble-Score
func Score(str string) (x int) {
	str = strings.ToUpper(str)
	Score := map[string]int{
		"A, E, I, O, U, L, N, R, S, T": 1,
		"D, G":                         2,
		"B, C, M, P":                   3,
		"F, H, V, W, Y":                4,
		"K":                            5,
		"J, X":                         8,
		"Q, Z":                         10,
	}
	for _, s := range str {
		for k, v := range Score {
			if strings.ContainsAny(k, string(s)) {
				x += v
			}
		}
	}
	return x
}
