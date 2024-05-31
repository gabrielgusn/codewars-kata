package main

import "fmt"

func IsPrime(n int) bool {

	var multipliers_count int = 0

	if n < 0 || n == 0 {
		return false
	}

	var index = 1

	for index <= n {
		if n%index == 0 {
			multipliers_count++
			if index < n && multipliers_count == 2 {
				return false
			}
		}
	}

	return multipliers_count == 2
}

func main() {
	fmt.Println(IsPrime(4294967295))
}
