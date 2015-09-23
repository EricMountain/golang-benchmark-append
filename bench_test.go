// Benchmarks comparing append() to straight copy
//
// Run:
//   for i in $(seq 1000 1000 10000000) ; do echo -- Size $i ; go test -bench . -- $i ; done > results.txt 2>&1
//
// With memory benchmarks:
//   for i in $(seq 1000 1000 10000000) ; do echo -- Size $i ; go test -bench . -benchmem -- $i ; done > results.txt 2>&1
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
	os.Exit(m.Run())
}

// TestDummy avoids runtime warning about no tests
func TestDummy(t *testing.T) {
}

func bDoIdx(d int, b *testing.B) {
	setup(d)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		doIdx(d)
	}
}

func bDoAppend(d int, b *testing.B) {
	setup(d)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		doAppend()
	}
}

func bDoPolo(d int, b *testing.B) {
	psetup(d)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		doPolo()
	}
}

func bDoPoloAppend(d int, b *testing.B) {
	psetup(d)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		doPoloAppend()
	}
}

func BenchmarkDoIdx(b *testing.B) {
	bDoIdx(dimension, b)
}

func BenchmarkDoAppend(b *testing.B) {
	bDoAppend(dimension, b)
}
func BenchmarkPolo(b *testing.B) {
	bDoPolo(dimension, b)
}

func BenchmarkPoloAppend(b *testing.B) {
	bDoPoloAppend(dimension, b)
}
