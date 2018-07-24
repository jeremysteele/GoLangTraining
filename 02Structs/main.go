package main

import "fmt"

type Person struct {
	fname string
	lname string
	age int
}

func (p Person) fullName() string {
	return p.fname + " " + p.lname
}

func main() {
	p1 := Person{"Jeremy", "Steele", 362}

	fmt.Println(p1.fullName())
}