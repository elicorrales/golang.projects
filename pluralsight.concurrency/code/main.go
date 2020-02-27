package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var maxBooks = 200
var sleepTime time.Duration = 80
var doConcurrently = false

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var lenOfCache = 0

func main() {

	configureAccordingToParams()

	queryCacheAndDatabaseLoop()
}

func configureAccordingToParams() {

	loops, mbErr := strconv.Atoi(os.Args[1])
	sleep, slErr := strconv.Atoi(os.Args[2])
	if len(os.Args) > 3 && os.Args[3] == "m" {
		doConcurrently = true
	}
	println(loops)
	fmt.Printf("loops err: %s\n", mbErr)
	println(sleep)
	fmt.Printf("loop sleep err: %s\n", slErr)

	if mbErr != nil || slErr != nil {
		os.Exit(1)
	}

	maxBooks = loops
	sleepTime = time.Duration(sleep)

}

func queryCacheAndDatabaseLoop() {

	start := time.Now().UnixNano()

	if doConcurrently {
		go kickOffQueryGoRoutines()
	} else {
		kickOffQueryGoRoutines()
	}

	//wait on goroutines to finish
	for lenOfCache < len(books) {
		println("Waiting on cache to fill")
		time.Sleep(100 * time.Millisecond)
	}

	end := time.Now().UnixNano()

	delta := end - start

	println("")
	println("------------------------")
	println(delta / 1000000)

}

func kickOffQueryGoRoutines() {

	for i := 0; i < maxBooks; i++ {

		print(lenOfCache)
		print(" ")

		if lenOfCache == len(books) {
			println("All books in cache")
			break
		}

		id := rnd.Intn(len(books)) + 1
		//id := i + 1

		if doConcurrently {
			go queryCache(id)
			go queryDatabase(id)
		} else {
			queryCache(id)
			queryDatabase(id)
		}

		// if we dont have this, then all threads(goroutines)
		// are quickly started, then this loop is done,
		// then main quits before all the goroutines
		// have a chance to run
		time.Sleep(sleepTime * time.Millisecond)
	}

}

// this query's purpose really is just to track num books in cache
func queryCache(id int) {
	if _, ok := cache[id]; ok {
		lenOfCache = len(cache)
	}
}

// if book is found in database, it is automatically added to cache.
// since it's map, we can re-add without any effects..no need to test.
// based on key.
func queryDatabase(id int) {

	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			//add found book to cache
			cache[id] = b
			return
		}
	}
}
