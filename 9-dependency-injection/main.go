package main

import (
	"log"
	"net/http"

	. "github.com/louischering/learn-go-with-tests/9-dependency-injection/dependencyinjection"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
