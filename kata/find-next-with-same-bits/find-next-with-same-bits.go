// package kata
package main

import (
	"fmt"
	"strconv"
)

func CountBits(bit_representation string, max_count int) int {
	count := 0

	for _, char := range bit_representation {
		if max_count > 0 && count > max_count {
			fmt.Println(bit_representation, "Brekou no", count, "maior que", max_count)
			return count
		}
		if string(char) == string("1") {
			count++
		}
	}
	return count
}

func NextHigher(n int) int {
	found_number := false
	n_bit_representation := strconv.FormatInt(int64(n), 2)
	n_bits := 0
	only_first_bits_are_1 := true

	for pos, char := range n_bit_representation {
		if string(char) == string("1") {
			if pos == n_bits {
				n_bits++
			} else {
				n_bits++
				only_first_bits_are_1 = false
			}
		}
	}

	aux_n := n + 1

	if only_first_bits_are_1 {
		aux_bit_string := "1"
		for i := 0; i < len(n_bit_representation); i++ {
			aux_bit_string += "0"
		}
		next_num, _ := strconv.ParseInt(aux_bit_string, 2, 32)
		fmt.Println(n, "que eh", n_bit_representation, "tem so 1 no comeso e retornou", next_num, "que eh", aux_bit_string)
		aux_n = int(next_num)
	}
	fmt.Println(n, "que eh", n_bit_representation)
	fmt.Println(aux_n, "que eh", strconv.FormatInt(int64(aux_n), 2))
	for !found_number {

		if CountBits(strconv.FormatInt(int64(aux_n), 2), n_bits) == n_bits {
			break
		}
		aux_n++
	}
	return aux_n
}

func main() {
	teste := NextHigher(805306367)
	fmt.Println(teste)
	fmt.Println(strconv.FormatInt(int64(teste), 2), 0)
}
