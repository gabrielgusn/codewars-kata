package main

import "fmt"

func Quadratic(x1, x2 int) [3]int {

	x1 *= -1
	x2 *= -1

	var coefficients_array = [3]int{}

	coefficients_array[0] = 1

	var second_coefficient int = x1 + x2

	coefficients_array[1] = second_coefficient

	var third_coefficient int = x1 * x2

	coefficients_array[2] = third_coefficient

	return coefficients_array
}

func main() {
	fmt.Printf("\n\n %v", Quadratic(4, 9))
}
