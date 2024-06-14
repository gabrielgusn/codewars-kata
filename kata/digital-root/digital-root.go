package main

import (
	"fmt"
	"strconv"
)

func DigitalRoot(n int) int {
	// if n < 10 {
	// 	return n
	// }
	var str string = strconv.Itoa(n)
	var sum int = 0
	for n < 10 {
		for i := 0; i < len(str); i++ {
			fmt.Println(str[i])
			sum += int(str[i])
		}
		n = sum
	}

	return n
	// return DigitalRoot(sum)
}

func main() {
	fmt.Println(DigitalRoot(16))
}
