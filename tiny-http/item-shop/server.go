package main

import (
	"log"
	"net/http"
)

func main() {
	port := "3000"

	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
