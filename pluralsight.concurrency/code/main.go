package main

import (
	//"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

var maxBooks = 200
var sleepTime time.Duration = 80
var doConcurrently = false
var doWriteLock = false
var doReadLock = false

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

	if len(os.Args) > 4 && os.Args[4] == "y" {
		doWriteLock = true
	}

	if len(os.Args) > 5 && os.Args[5] == "y" {
		doReadLock = true
	}

	//println(loops)
	//fmt.Printf("loops err: %s\n", mbErr)
	//println(sleep)
	//fmt.Printf("loop sleep err: %s\n", slErr)

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
		println("")
		println("_W_")
		time.Sleep(20000 * time.Microsecond)
		//if (time.Now().UnixNano()-start)/1000000 > 1000 {
		if (time.Now().UnixNano()-start)/1000 > 1000000 {
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

func kickOffQueryGoRoutines() {

	for i := 0; i < maxBooks; i++ {

		print(lenOfCache)
		print(" ")

		if lenOfCache == len(books) {
			println("")
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
		time.Sleep(sleepTime * time.Microsecond)
	}

}

// this query's purpose really is just to track num books in cache
func queryCache(id int) {
	if doReadLock {
		mutex.Lock()
	}
	if _, ok := cache[id]; ok {
		lenOfCache = len(cache)
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
