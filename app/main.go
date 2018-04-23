package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// initialize data
type person struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *address `json:"address,omitempty"`
}
type address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

var people []person

//GET method
// homePage
func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	welcomeMessage := `
	<div style="display:grid;height:75vh">
		<div style="margin:auto">
			<h1 style="text-align:center; font-size:4em"> Welcome to SQL ANGELES </h1>
			<h1 style="text-align:center; font-size:2em"> This project is a simple RESTful API using Golang with Docker </h1>
			<div style="display:grid;height:100%">
			<ul style="margin: auto;list-style:none; font-size: 1.7em">
				<li>localhost:8000 - GET - Home Page</li>
				<li>localhost:8000/people - GET - getPeople</li>
				<li>localhost:8000/people/:id - GET - getPerson</li>
				<li>localhost:8000/people/:id - POST - createPerson</li>
				<li>localhost:8000/people/:id - DELETE - deletePerson</li>
			</ul>
			</div>
		</div>
	</div>
	`
	fmt.Fprintf(w, welcomeMessage)
}

// GET method
// get all people
func getPeople(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(people)
}

// GET method
// get one person by passing id
func getPerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	personID := ps.ByName("id")
	for _, item := range people {
		if item.ID == personID {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Fprintf(w, "<h1>No DATA</h1>")
}

// POST method
// create person by passing id in url and dat in json
func createPerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newID := ps.ByName("id")
	var newPerson person
	_ = json.NewDecoder(r.Body).Decode(&newPerson)
	newPerson.ID = string(newID)
	people = append(people, newPerson)
	json.NewEncoder(w).Encode(people)
}

// DELET method
// delete person by passing id
func deletePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	personID := ps.ByName("id")
	for index, item := range people {
		if item.ID == personID {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

// main func
func main() {
	// initialize router
	router := httprouter.New()

	// add data
	people = append(people, person{ID: "1", Name: "Sangwoo", Age: 29, Address: &address{City: "Los Angeles", State: "CA"}})
	people = append(people, person{ID: "2", Name: "Paul", Age: 28, Address: &address{City: "Irvine", State: "CA"}})

	// routers
	router.GET("/", homePage)
	router.GET("/people", getPeople)
	router.GET("/people/:id", getPerson)
	router.POST("/people/:id", createPerson)
	router.DELETE("/people/:id", deletePerson)

	// prints the message on bash
	log.Println("running api server on port 8000")

	// listens port 8000 and add router
	http.ListenAndServe(":8000", router)
}
