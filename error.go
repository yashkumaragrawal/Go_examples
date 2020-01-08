package main

import (
	"fmt"
	"math"
)
type MyError struct {
	a string	
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v",e.a)
}
func Sqrt(x float64) (float64, error) {
    if x > 0 {
		return math.Sqrt(x), nil
	}
	return 0, &MyError{"error"}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
