package main

import (
  "fmt"
	"net/http"
	"log"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
  // Parse arguments
  r.ParseForm()

  // Print form info
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
    fmt.Println("key:", k)
    fmt.Println("val:", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello, Erick")
}

func main() {
  // Set Router
  http.HandleFunc("/", sayHelloName)

  // Set listen port
  err := http.ListenAndServe(":9090", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
