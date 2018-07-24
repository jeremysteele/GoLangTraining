package main

import "fmt"

func madness(ptr **int) {
	**ptr++
}

func supermadness(ptr ***int) {
	***ptr++
}

func main() {
	a := 5
	var  b *int = &a
	c := &b

	fmt.Println("Orig")
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(*c)
	fmt.Println(**c)

	**c = 10

	fmt.Println("\n\nNew")
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(*c)
	fmt.Println(**c)

	fmt.Println("\n\nMADNESS")
	*b = 1
	fmt.Println(*b)
	madness(&b)
	fmt.Println(*b)

	fmt.Println("\n\nSUPER MADNESS")
	**c = 1
	fmt.Println(**c)
	supermadness(&c)
	fmt.Println(**c)
}
