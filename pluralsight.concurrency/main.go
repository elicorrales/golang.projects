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
var doWriteLock = false
var doReadLock = false

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var lenOfCache = 0

func main() {

	commandLineHandler()

	start := time.Now().UnixNano()

	for i := 0; i < maxBooks; i++ {

		if lenOfCache == len(books) {
			println("")
			println("All books in cache")
			break
		}

		id := rnd.Intn(len(books)) + 1
		if doConcurrently {
			go queryCache(id)
			go queryDatabase(id)
		} else {
			queryCache(id)
			queryDatabase(id)
		}

	}

	//wait on goroutines to finish
	for lenOfCache < len(books) {
		//println("")
		//println("_W_")
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

	if len(os.Args) < 7 {
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

	switch os.Args[6] {
	case "y":
		doWriteLock = true
	case "n":
		doWriteLock = false
	default:
		{
			usage()
		}
	}

	switch os.Args[7] {
	case "y":
		doReadLock = true
	case "n":
		doReadLock = false
	default:
		{
			usage()
		}
	}

	/*


		maxBooks = loops
		sleepTime = time.Duration(sleep)
	*/
}

func usage() {
	println("")
	println("")
	println("prog <books> <loops> <sleep> <random y|n> <threaded y|n> <y> <y>")
	println("")
	println("")
	os.Exit(1)
}

// this query's purpose really is just to track num books in cache
func queryCache(id int) {

	print(".")

	if doReadLock {
		mutex.Lock()
	}
	_, ok := cache[id]
	if ok {
		fmt.Printf("%d ", len(cache))
	}
	if doReadLock {
		mutex.Unlock()
	}
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
			if doWriteLock {
				mutex.Lock()
			}
			cache[id] = b
			if doWriteLock {
				mutex.Unlock()
			}
			return
		}
	}
}
