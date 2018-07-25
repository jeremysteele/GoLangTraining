package main

import (
	"fmt"
	"math/rand"
	"time"
	"regexp"
	"log"
	"strconv"
)

func rollDice(sides int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(sides) + 1
}

func main() {
	var roll string

	fmt.Println("Enter your roll in 3d6 format, where 3 is the number of dice and 6 is the number of sides: ")
	fmt.Scanln(&roll)

	r, _ := regexp.Compile("^([\\d]+)d([\\d]+)$")

	roll_match := r.FindStringSubmatch(roll)

	if roll_match == nil {
		fmt.Println("Bad input")
		log.Fatal("Incorrect input: ", roll)
	}

	numRolls, _ := strconv.ParseInt(roll_match[1], 10, 64)
	numSides, _ := strconv.ParseInt(roll_match[2], 10, 64)
	output := ""
	curTotal := 0

	for i:= int64(0); i < numRolls; i++ {
		curRoll := rollDice(int(numSides))
		curTotal += curRoll

		output += fmt.Sprintf("%d ", curRoll)
	}

	fmt.Printf("%d: %s\n", curTotal, output)
}
