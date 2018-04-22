package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!! My first Golang")
}

func main() {

	http.HandleFunc("/", HomeEndPoint)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
