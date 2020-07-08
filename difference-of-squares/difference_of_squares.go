package diffsquares

import "math"

//the function SquareOsSum returns the square of the sum of n digits
func SquareOfSum(x int) int {
	var sum = 0.0
	for i := float64(1); i <= float64(x); i++ {
		sum += i
	}
	return int(math.Pow(sum, 2))
}

//the function SumOfSquare returns the square of the sum of n numbers
func SumOfSquares(x int) int {
	var sqr = 0.0
	for i := float64(1); i <= float64(x); i++ {
		sqr += math.Pow(i, 2)
	}
	return int(sqr)
}

//the function Difference finds the difference of two function results
func Difference(x int) int {
	return (SquareOfSum(x) - SumOfSquares(x))
}
