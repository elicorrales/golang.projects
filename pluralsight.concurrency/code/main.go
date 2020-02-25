package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxBooks = 50

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var lenOfCache = 0

func main() {

	queryCacheAndDatabaseLoop()

}

func queryCacheAndDatabaseLoop() {

	start := time.Now().UnixNano()

	for i := 0; i < maxBooks; i++ {

		if lenOfCache == 10 {
			println("All books in cache")
			break
		}

		id := rnd.Intn(10) + 1
		//id := i + 1

		queryCacheAndDatabase(id)

		//gives cache time to receive new data in concurrent operation (???)
		time.Sleep(120 * time.Millisecond)
	}

	end := time.Now().UnixNano()

	delta := end - start

	println("------------------------")
	println(delta / 1000000)

}

func queryCacheAndDatabase(id int) {

	go func(id int) {
		if _, ok := queryCache(id); ok {
			fmt.Printf("In cache : %d, %t, %+v\n", len(cache))
			lenOfCache = len(cache)
		}
	}(id)

	go queryDatabase(id)
	//queryDatabase(id)

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
