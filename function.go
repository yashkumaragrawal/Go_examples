package main
import "fmt"
func fibonacci(n int) {
	x := 0
	y := 1
	for l := 0; l < n; l++ {
		v := x
		fmt.Println(v)
		x = y
		y = y + v
	}
}
func main() {
    	fibonacci(10)
}

