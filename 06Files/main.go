package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	fh, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := fh.Read(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes %q\n", count, data[:count])
}
