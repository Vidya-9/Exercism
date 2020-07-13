package grains

import (
	"errors"
	"fmt"
)

//definition of the function Square
func Square(n int) (uint64, error) {

	fmt.Scan(&n)
	if n <= 0 || n > 64 {
		return 0, errors.New("true")
	}
	return 1 << (n - 1), nil
	//func pow(x int)uint64{
	//x = 1
	//for i := 1; i <= n; i++ {
	//Square = x **x
	//}
	//}
	//return x, errors.New("")
}

//the definition of the function Total
func Total() uint64 {
	return 1<<64 - 1
}
