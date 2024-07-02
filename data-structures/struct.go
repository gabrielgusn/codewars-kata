package main

import "fmt"

type contactInfo struct {
	phone string
	mail  string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	gabriel := person{
		firstName: "Gabriel",
		lastName:  "Nicolodi",
		contactInfo: contactInfo{
			phone: "4699999999",
			mail:  "gabriel@gabriel.com",
		},
	}
	gabriel.updateName("teste")
	gabriel.print()
}

func (p *person) updateName(newName string) {
	p.firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}
