package main

import "fmt"

func main() {
	fmt.Println(convertDecimalToBinary(0))
	fmt.Println(CountBits(0))
}

func CountBits(num uint) int {
	var bitArray []int8 = convertDecimalToBinary(num)

	var count int = 0

	for i := 0; i < len(bitArray); i++ {
		if bitArray[i] == 1 {
			count++
		}
	}

	return count
}

func convertDecimalToBinary(num uint) []int8 {
	var array []int8
	if num == 0 {
		return append(array, 0)
	}

	for num > 1 {
		array = prepend(array, int8(num%2))
		num /= 2
	}
	array = prepend(array, int8(1))
	return array
}

func prepend(slice []int8, element int8) []int8 {
	return append([]int8{element}, slice...)
}
