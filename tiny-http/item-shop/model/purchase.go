package main

import (
	"sync"
)

// Our "Database" file
const purchasesFile = "db/purchases.json"

// A Purchase is what we get when
// we make an order of sorts.
//
// We will auto generate the id, items
// are whatever the client passes, and our total is
// summed up from the Item array
type Purchase struct {
	id    int
	items []Item
	total int
}

// This, is a "Mutural exclusion lock" which allows us to
// lock our files so that we don't access it twice at the same time
// and do some bad stuff!
var purchasesMutex = new(sync.Mutex)

// GetByID lets us grab a specific purchase using its ID
func (purchase *Purchase) GetByID(id int) (Purchase, error) {
	return nil, nil
}

// GetAll grabs every purchase in our database and returns it
func (purchase *Purchase) GetAll() ([]Purchase, error) {
	return nil, nil
}

// Create saves a purchase to the database, and returns it
func (purchase *Purchase) Create(p Purchase) (Purchase, error) {
	return nil, nil
}
