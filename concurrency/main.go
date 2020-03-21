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

var start = time.Now().UnixNano()

//================================================================
func main() {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	commandLineHandler()

	id := 0
	for i := 0; i < maxLoops; i++ {

		if doRandom {
			id = rnd.Intn(maxBooks) + 1
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
	println("Main is waiting....")
	time.Sleep(1000 * time.Millisecond)
	end := time.Now().UnixNano()
	delta := end - start
	println("")
	println("Main is done.")
	println("")
	print(delta / 1000000)
	println(" ms")
	println("")

}

//================================================================
func commandLineHandler() {

	if len(os.Args) < 6 {
		usage()
	}

	num_books, mbErr := strconv.Atoi(os.Args[1])
	num_loops, mlErr := strconv.Atoi(os.Args[2])
	sleep, slErr := strconv.Atoi(os.Args[3])

	if mbErr != nil || mlErr != nil || slErr != nil {
		usage()
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

	if num_books < 1 {
		println("")
		println("Num books must be at least 1")
		usage()
	}

	if num_books > len(books) {
		println("")
		fmt.Printf("Num books must be no more than %d\n", len(books))
		usage()
	}

	maxBooks = num_books
	maxLoops = num_loops
	sleepTime = time.Duration(sleep)

}

//================================================================
func usage() {
	println("")
	println("")
	println("prog <books> <loops> <sleep> <random y|n> <threaded y|n> ")
	println("")
	println("")
	os.Exit(1)
}

//================================================================
func exitWhenFull() {
	println("")
	println("")
	println("Cache is full.")
	println("")
	end := time.Now().UnixNano()
	delta := end - start
	print(delta / 1000000)
	println(" ms")
	println("")
	os.Exit(0)
}

//================================================================
func queryCache(id int) {

	print(".")

	mutex.Lock()

	_, ok := cache[id]
	if ok {
		fmt.Printf("%d ", len(cache))
	}

	mutex.Unlock()

	if len(cache) >= maxBooks {
		exitWhenFull()
	}
}

//================================================================
func queryDatabase(rndId int) {

	time.Sleep(sleepTime * time.Millisecond)
	for _, b := range books {
		if b.ID == rndId {
			mutex.Lock()
			cache[rndId] = b
			mutex.Unlock()
			break
		}
	}
}
