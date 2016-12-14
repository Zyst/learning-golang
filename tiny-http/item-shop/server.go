package main

import (
	"log"
	"net/http"
	"sync"
)

const itemsFile = "db/items.json"
const purchasesFile = "db/purchases.json"

// This is used so we can edit files without worrying about
// multiple API calls
var itemsMutex = new(sync.Mutex)
var purchasesMutex = new(sync.Mutex)

func main() {
	port := "3000"

	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
