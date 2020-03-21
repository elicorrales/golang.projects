package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Date   string
}

var cache = map[int]Book{}

func main() {
	fmt.Printf("map len:%d\n", len(cache))
	cache[1] = Book{}
	fmt.Printf("map len:%d\n", len(cache))
	cache[2] = Book{}
	fmt.Printf("map len:%d\n", len(cache))
}
