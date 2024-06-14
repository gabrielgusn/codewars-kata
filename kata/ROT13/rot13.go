package main

import "fmt"

// import "strconv"

func Rot13(msg string) string {
	var cypher string
	fmt.Println(rune('0'))
	for pos, char := range msg {

		// fmt.Println(char)
		if char >= rune('a') && char <= rune('z') {
			if char <= rune('m') {
				cypher += string(char + 13)
			} else {
				cypher += string((rune('z') - char) + rune('a'))
			}
		} else if char >= rune('A') && char <= rune('Z') {
			if char <= rune('M') {
				cypher += string(char + 13)
			} else {
				cypher += string((rune('Z') - char) + rune('A'))
			}
		} else if char >= rune('0') && char <= rune('9') {
			fmt.Println("NUMERO", char, string(char))
			cypher += string(char)
		} else if pos > 0 && rune(msg[pos-1]) == rune('\\') {
			continue
		} else {
			continue
		}
	}

	return cypher
}

func main() {
	fmt.Println(Rot13("EBG13 rknzcyr."))
}
