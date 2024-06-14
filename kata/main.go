package main

import "fmt"

func main() {

	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	fmt.Print("x ", x, "\n")
	fmt.Print("ip ", ip, "\n")
	fmt.Print("y ", y, "\n")

	*ip = 2

	fmt.Print("x ", x, "\n")
	fmt.Print("ip ", ip, "\n")
	fmt.Print("y ", y, "\n")

}
