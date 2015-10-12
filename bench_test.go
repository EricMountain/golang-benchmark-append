// Benchmarks comparing append() to straight copy
//
// Run:
//   for i in $(seq 10000 10000 10000000) ; do echo -- Size $i ; go test -bench . -- $i ; done > results.txt 2>&1
//
// With memory benchmarks:
//   for i in $(seq 10000 10000 10000000) ; do echo -- Size $i ; go test -bench . -benchmem -- $i ; done > results.txt 2>&1
//
// How long it ran:
// grep ok results.txt | awk '{print($3)}' | sed -e 's/s//' | awk 'BEGIN { x = 0 } {x+=$1} END {print(x)}'
//
// How many tests ran:
// grep ok results.txt | wc -l
//
package main

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
)

type summat struct {
	n string
	i int
}

var (
	dimension = 10
	source    []int
	psource   map[int]summat
)

func TestMain(m *testing.M) {
	// Yuck.  Size of the data to use for the struct is last arg.
	if d, err := strconv.Atoi(os.Args[len(os.Args)-1]); err == nil {
		dimension = d
	}

	setup(dimension)

	os.Exit(m.Run())
}

// TestDummy avoids runtime warning about no tests
func TestDummy(t *testing.T) {
}

func BenchmarkIntIdx(b *testing.B) {
	bench(b, doIntIdx)
}

func BenchmarkIntAppend(b *testing.B) {
	bench(b, doIntAppend)
}

func BenchmarkPoloIdx(b *testing.B) {
	bench(b, doPoloIdx)
}

func BenchmarkPoloAppend(b *testing.B) {
	bench(b, doPoloAppend)
}

func bench(b *testing.B, f func()) {
	for n := 0; n < b.N; n++ {
		f()
	}
}

func setup(dim int) {
	source = make([]int, dim)
	for i := range source {
		source[i] = rand.Int()
	}

	psource = make(map[int]summat, dim)
	for i := 0; i < dim; i++ {
		psource[i] = summat{n: strconv.Itoa(rand.Int()), i: rand.Int()}
	}
}

func doIntIdx() {
	dest := make([]int, len(source))
	for i := range source {
		dest[i] = source[i]
	}
}

func doIntAppend() {
	dest := make([]int, 0, len(source))
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
func doPoloIdx() {
	dest := make([]summat, len(psource))
	i := 0
	for _, v := range psource {
		dest[i] = v
		i++
	}
}

// Polo's case using append()
func doPoloAppend() {
	dest := make([]summat, 0, len(psource))
	for _, v := range psource {
		dest = append(dest, v)
	}
}
