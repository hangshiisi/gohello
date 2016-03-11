package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	RunPolicyAgent()

	log.Fatal(http.ListenAndServe(":8080", router))
}
