package main

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	FirstName string
	LastName string
}

type APIError struct {
	Message string
	Code int
}

var session *mgo.Session

func personEndpoint(w http.ResponseWriter, r *http.Request) {
	switch(r.Method) {
	case "GET":
		getPerson(w, r)
	case "POST":
		addPerson(w, r)
	}
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || len(r.URL.Query()) < 1 {
		json.NewEncoder(w).Encode(APIError{"Invalid Request", 500})
		return
	}

	result := Person{}
	c := session.DB("test").C("people")
	err := c.Find(bson.M{"firstname": r.URL.Query().Get("firstname")}).One(&result)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(APIError{"Not Found", 404})
		return
	}

	json.NewEncoder(w).Encode(result)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		json.NewEncoder(w).Encode(APIError{"Invalid Request", 500})
		return
	}

	decoder := json.NewDecoder(r.Body)
	var p Person

	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}

	c := session.DB("test").C("people")
	err = c.Insert(&p)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(APIError{"Success", 200})
}

func connect() {
	sess, err := mgo.Dial("mongodb://root:password@mongo:27017")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	sess.SetMode(mgo.Monotonic, true)
	session = sess
}

func main() {
	connect()
	defer session.Close()

	http.HandleFunc("/person", personEndpoint)

	log.Println("Starting server http://localhost:8080/")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
