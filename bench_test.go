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
	"os"
	"strconv"
	"testing"
)

var (
	dimension = 10
)

func TestMain(m *testing.M) {
	// Yuck.  Size of the data to use for the struct is last arg.
	if d, err := strconv.Atoi(os.Args[len(os.Args)-1]); err == nil {
		dimension = d
	}

	setup(dimension)
	psetup(dimension)

	os.Exit(m.Run())
}

// TestDummy avoids runtime warning about no tests
func TestDummy(t *testing.T) {
}

func bench(b *testing.B, f func()) {
	for n := 0; n < b.N; n++ {
		f()
	}
}

func BenchmarkIntIdx(b *testing.B) {
	bench(b, doIdx)
}

func BenchmarkIntAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doAppend()
	}
}

func BenchmarkPoloIdx(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doPolo()
	}
}

func BenchmarkPoloAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		doPoloAppend()
	}
}
