package main

import (
	"fmt"
	"math/rand"
)

var (
	source []int
)

func main() {
	dim := 10
	setup(dim)
	for i := range source {
		fmt.Println(source[i])
	}
	doIdx(dim)
}

func setup(dim int) {
	source = make([]int, dim)
	for i := range source {
		source[i] = rand.Int()
	}
}

func doIdx(dim int) {
	var dest []int
	dest = make([]int, dim)
	for i := range source {
		dest[i] = source[i]
	}
}

func doAppend(dim int) {
	var dest []int
	dest = make([]int, 0, dim) // setting initial length to something other than 0 is a bad idea for append()
	for i := range source {
		dest = append(dest, source[i])
	}
}
