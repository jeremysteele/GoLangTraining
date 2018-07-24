package main

import (
	"errors"
	"fmt"
)

func bad(a int) (bool, error) {
	if(a == 0) {
		return false, errors.New("Derp")
	}

	return true, nil
}

func main() {
	a, err := bad(5)

	fmt.Println(a)
	fmt.Println(err)

	a, err = bad(0)

	fmt.Println(a)
	fmt.Println(err)
}
