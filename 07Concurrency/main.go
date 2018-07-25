package main

import (
	"fmt"
	"time"
)

func firstThing() {
	for x := 0; x < 5;x++ {
		time.Sleep(500)
		fmt.Printf("x: %d\n", x)
	}
}

func secondThing() {
	for y := 0; y < 5;y++ {
		time.Sleep(500)
		fmt.Printf("y: %d\n", y)
	}
}

func main() {
	go firstThing()
	go secondThing()

	for true {}
}
