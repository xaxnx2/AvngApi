package main

import (
	"HOMEWORKAVENGINVAPI/Internal/Router"
	"log"
	"net/http"
)

func main() {
	r := Router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
