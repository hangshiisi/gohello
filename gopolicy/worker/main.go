package main

import (
	"github.com/hangshiisi/gohello/gopolicy/worker/rest"
	"github.com/hangshiisi/gohello/gopolicy/worker/scheduler"
	"log"
	"net/http"
)

func main() {

	router := rest.NewRouter()

	scheduler.RunPolicyAgent()

	log.Fatal(http.ListenAndServe(":8080", router))
}
