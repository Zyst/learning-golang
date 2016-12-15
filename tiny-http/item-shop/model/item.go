package main

import (
	"sync"
)

// Our "Database" file!
const itemsFile = "../db/items.json"

// Item is what we get from our .json file
// and which we handle as "Store Items" so
// to say. All of our operations should be "GET"
// types, we don't really delete Items at any point
type Item struct {
  id int
  name string
  price int
}

// This, is a "Mutural exclusion lock" which allows us to
// lock our files so that we don't access it twice at the same time
// and do some bad stuff!
var itemsMutex = new(sync.Mutex)

// GetByID returns the item bound to the id passed
// to the function
func (item *Item) GetByID(id int) (Item, error) {
  return nil, nil
}

// GetAll returns every item in our shop in an array,
// and an error if something went horribly wrong
func (item *Item) GetAll() ([]Item, error) {
  return nil, nil
}
