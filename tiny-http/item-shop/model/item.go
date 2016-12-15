package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"
)

// Our "Database" file
const itemsFile = "../db/items.json"

// Item is what we get from our .json file
// and which we handle as "Store Items" so
// to say. All of our operations should be "GET"
// types, we don't really delete Items at any point
type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// This, is a "Mutural exclusion lock" which allows us to
// lock our files so that we don't access it twice at the same time
// and do some bad stuff!
var itemsMutex = new(sync.Mutex)

// GetByID returns the item bound to the id passed
// to the function
func (item *Item) GetByID(id int) (Item, error) {
	// We lock when doing file operations
	itemsMutex.Lock()

	// When the function is done executing we unlock
	defer itemsMutex.Unlock()

	itemData, err := ioutil.ReadFile(itemsFile)
	if err != nil {
		return Item{}, err
	}

	var items []Item

	// Here while doing the json Unmarshal we pass a pointer of
	// our items array, so items is now filled
	if err := json.Unmarshal(itemData, &items); err != nil {
		return Item{}, err
	}

	for _, itm := range items {
		if itm.ID == id {
			return itm, nil
		}
	}

	return Item{}, errors.New("I cannot give you the item with that ID traveler," +
		" my strongest items would kill you")
}

// GetAll returns every item in our shop in an array,
// and an error if something went horribly wrong
func (item *Item) GetAll() ([]Item, error) {
	// We lock when doing file operations
	itemsMutex.Lock()

	// When the function is done executing we unlock
	defer itemsMutex.Unlock()

	itemData, err := ioutil.ReadFile(itemsFile)
	if err != nil {
		return nil, err
	}

	var items []Item

	// Here while doing the json Unmarshal we pass a pointer of
	// our items array, so items is now filled
	if err := json.Unmarshal(itemData, &items); err != nil {
		return nil, err
	}

	return items, nil
}
