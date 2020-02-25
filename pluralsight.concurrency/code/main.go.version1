package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxBooks = 10

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	queryCacheAndDatabaseLoop()

	queryCacheAndDatabaseLoop()

}

func queryCacheAndDatabaseLoop() {

	cacheCount := 0
	dbCount := 0

	for i := 0; i < maxBooks; i++ {

		id := rnd.Intn(maxBooks) + 1
		//id := i + 1

		if b, ok := queryCache(id); ok {

			fmt.Printf("In cache : %d, %t, %+v\n", id, ok, b)
			cacheCount++
			continue

		}

		if b, ok := queryDatabase(id); ok {

			fmt.Printf("In DB: %d, %t, %+v\n", id, ok, b)
			dbCount++
			continue
		}

		fmt.Printf("================ %d not found.\n", id)

		//gives cache time to receive new data in concurrent operation (???)
		time.Sleep(150 * time.Millisecond)

	}

	fmt.Printf("Cache hits: %d , DB hits: %d ===========================\n", cacheCount, dbCount)

}

func queryCache(id int) (Book, bool) {

	b, ok := cache[id]
	return b, ok

}

func queryDatabase(id int) (Book, bool) {

	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			//add found book to cache
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
