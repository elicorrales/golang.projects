package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

var maxBooks = 0
var maxLoops = 0
var sleepTime time.Duration = 0
var doRandom = false
var doConcurrently = false

var cache = map[int]Book{}

//=================================================================================================
func main() {

	start := time.Now().UnixNano()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	commandLineHandler()

	id := 0
	for i := 0; i < maxLoops; i++ {

		if doRandom {
			id := rnd.Intn(len(maxBooks)) + 1
		} else {
			if id >= maxBooks {
				id = 1
			} else {
				id++
			}
		}

		if doConcurrently {
			go queryCache(id)
			go queryDatabase(id)
		} else {
			queryCache(id)
			queryDatabase(id)
		}

	}

	println("")
	println("")
	timedOut := false
	start_wait := time.Now().UnixNano()
	for lenOfCache < len(books) {
		time.Sleep(20 * time.Millisecond)
		//time.Sleep(20000 * time.Microsecond)
		if (time.Now().UnixNano()-start)/1000000 > 1000 {
			//if (time.Now().UnixNano()-start)/1000 > 1000000 {
			println("")
			println("Timed out waiting for queries to complete.")
			print("Was able to load ")
			println(lenOfCache)
			break
		}
	}

	end := time.Now().UnixNano()

	delta := end - start

	println("")
	println("------------------------")
	print(delta / 1000000)
	println(" ms")
	print(delta / 1000)
	println(" us")

}

func commandLineHandler() {

	if len(os.Args) < 6 {
		usage()
	}

	books, mbErr := strconv.Atoi(os.Args[1])
	loops, mlErr := strconv.Atoi(os.Args[2])
	sleep, slErr := strconv.Atoi(os.Args[3])

	if mbErr != nil || mlErr != nil || slErr != nil {
		os.Exit(1)
	}

	switch os.Args[4] {
	case "y":
		doRandom = true
	case "n":
		doRandom = false
	default:
		{
			usage()
		}
	}

	switch os.Args[5] {
	case "y":
		doConcurrently = true
	case "n":
		doConcurrently = false
	default:
		{
			usage()
		}
	}

	maxBooks = books
	maxLoops = loops
	sleepTime = time.Duration(sleep)

}

func usage() {
	println("")
	println("")
	println("prog <books> <loops> <sleep> <random y|n> <threaded y|n> ")
	println("")
	println("")
	os.Exit(1)
}

// this query's purpose really is just to track num books in cache
func queryCache(id int) {

	print(".")

	mutex.Lock()

	_, ok := cache[id]
	if ok {
		fmt.Printf("%d ", len(cache))
	}

	mutex.Unlock()
}

// if book is found in database, it is automatically added to cache.
// since it's map, we can re-add without any effects..no need to test.
// based on key.
func queryDatabase(id int) {

	time.Sleep(80000 * time.Microsecond)
	for _, b := range books {
		if b.ID == id {

			//fmt.Printf("rid:%d b.ID:%d ->same, added\n", id, b.ID)
			//add found book to cache

			mutex.Lock()

			cache[id] = b

			mutex.Unlock()

			return
		}
	}
}
