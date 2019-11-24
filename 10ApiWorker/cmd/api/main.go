package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jeremysteele/GoLangTraining/10ApiWorker/internal"
	"github.com/jeremysteele/GoLangTraining/10ApiWorker/models"
)

func addPerson(w http.ResponseWriter, req *http.Request) {
	client := internal.GetRedisClient()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Got request %s\n", body)

	var p models.Person
	err = json.Unmarshal(body, &p)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	str, _ := json.Marshal(p)

	err = client.LPush("people", str).Err()
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "added person to queue\n")
}

func getPeople(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello get\n")
}

func main() {
	client := internal.GetRedisClient()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	fmt.Printf("Starting up...\n")

	http.HandleFunc("/add", addPerson)
	http.HandleFunc("/get", getPeople)

	http.ListenAndServe("0.0.0.0:8080", nil)
}
