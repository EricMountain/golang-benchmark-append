package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type summat struct {
	n string
	i int
}

var (
	source  []int
	psource map[int]summat
)

func main() {
	dim := 10
	psetup(dim)
	for k, v := range psource {
		fmt.Println(k, v)
	}
	doPoloAppend()
}

func setup(dim int) {
	source = make([]int, dim)
	for i := range source {
		source[i] = rand.Int()
	}
}

func psetup(dim int) {
	psource = make(map[int]summat, dim)
	for i := 0; i < dim; i++ {
		psource[i] = summat{n: strconv.Itoa(rand.Int()), i: rand.Int()}
	}
}

func doIdx() {
	dest := make([]int, len(source))
	for i := range source {
		dest[i] = source[i]
	}
}

func doAppend() {
	dest := make([]int, 0, len(source)) // setting initial length to something other than 0 is a bad idea for append()
	for i := range source {
		dest = append(dest, source[i])
	}
}

// Polo's case
// map[int]struct
// copy map values to a slice :
// slice := make([]struct, len(map))
// i := 0
// for _, v := range map {
//   slice[i] = v
//   i++
// }
func doPolo() {
	dest := make([]summat, len(psource))
	i := 0
	for _, v := range psource {
		dest[i] = v
		i++
	}
}

// Polo's case using append()
func doPoloAppend() []summat {
	dest := make([]summat, 0, len(psource))
	for _, v := range psource {
		dest = append(dest, v)
	}
	return dest
}
