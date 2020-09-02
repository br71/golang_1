package main

import "fmt"

type locationInfo struct {
	email   string
	zipCode int
}

type person struct {
	name     string
	lastname string
	location locationInfo
}

func (p person) print() {
	fmt.Println(p.name, p.lastname)
}

func (p *person) updateName(newName string) {
	(*p).name = newName
}

func main() {
	alex := person{
		name:     "Alex",
		lastname: "Anderson",
		location: locationInfo{
			email:   "amderson@gmail.com",
			zipCode: 5154,
		},
	}

	alexPointer := &alex

	alex.print()
	alex.name = "Jim"
	alex.print()
	alexPointer.updateName("Nexus")
	alex.print()
	alex.updateName("Nexus2")
	alex.print()

}
