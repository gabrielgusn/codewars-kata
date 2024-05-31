// package kata
package main

import "fmt"

func Solution(str string) []string {

	var aux_array []string

	for pos, char := range str {
		if pos == len(str)-1 {
			if len(str)%2 != 0 {
				aux_array = append(aux_array, string(char)+"_")
			}
		}
		if pos%2 != 0 && pos > 0 {
			aux_array = append(aux_array, string(str[pos-1])+string(char))
		}
	}

	return aux_array
}

func main() {
	fmt.Println(Solution(""))
}
