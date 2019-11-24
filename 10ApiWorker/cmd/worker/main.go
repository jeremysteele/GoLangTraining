package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jeremysteele/GoLangTraining/10ApiWorker/internal"
	"github.com/jeremysteele/GoLangTraining/10ApiWorker/models"
)

func processPerson(p models.Person) {
	client := internal.GetRedisClient()

	result := fmt.Sprintf("%s %s - Age %d", p.FirstName, p.LastName, p.Age)

	fmt.Printf("Got Result: %s\n", result)

	client.LPush("results", result)
}

func main() {
	client := internal.GetRedisClient()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	runLoop := true

	var p models.Person

	for runLoop {
		fmt.Printf("Checking queue\n")

		val, err := client.LPop("people").Result()
		if err == nil && len(val) > 0 {
			fmt.Printf("Popped %s\n", val)

			err := json.Unmarshal([]byte(val), &p)

			if err != nil {
				fmt.Printf("Couldn't unmarshal object. Uh oh\n")
				continue
			}

			go processPerson(p)

		} else {
			fmt.Printf("Nothing yet. Gonna goto bed now. Bye\n")
		}

		time.Sleep(1 * time.Second)
	}
}
