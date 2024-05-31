// package kata
package main

import (
	"fmt"
	"strings"
)

func FirstNonRepeating(str string) string {

	var first_non_repeated_char string = ""

	if len(str) == 1 {
		return string(str[0])
	}

	for pos, char := range str {
		for pos2, char2 := range str {
			if pos2 == pos {
				continue
			}
			if strings.ToLower(string(char2)) == strings.ToLower(string(char)) {
				break
			}
			if pos2 == len(str)-1 &&
				strings.ToLower(string(char2)) != strings.ToLower(string(char)) {
				first_non_repeated_char = string(char)
				break
			}
		}
		if first_non_repeated_char != "" {
			break
		}
	}

	return first_non_repeated_char
}

func main() {
	fmt.Println(FirstNonRepeating("a"))
}
