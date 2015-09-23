package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
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
		doAppend(d)
	}
}

func BenchmarkDoIdx_10(b *testing.B) {
	bDoIdx(10, b)
}

func BenchmarkDoAppend_10(b *testing.B) {
	bDoAppend(10, b)
}

func BenchmarkDoIdx_100(b *testing.B) {
	bDoIdx(100, b)
}

func BenchmarkDoAppend_100(b *testing.B) {
	bDoAppend(100, b)
}

func BenchmarkDoIdx_1000(b *testing.B) {
	bDoIdx(1000, b)
}

func BenchmarkDoAppend_1000(b *testing.B) {
	bDoAppend(1000, b)
}

func BenchmarkDoIdx_10000(b *testing.B) {
	bDoIdx(10000, b)
}

func BenchmarkDoAppend_10000(b *testing.B) {
	bDoAppend(10000, b)
}

func BenchmarkDoIdx_100000(b *testing.B) {
	bDoIdx(100000, b)
}

func BenchmarkDoAppend_100000(b *testing.B) {
	bDoAppend(100000, b)
}

func BenchmarkDoIdx_1000000(b *testing.B) {
	bDoIdx(1000000, b)
}

func BenchmarkDoAppend_1000000(b *testing.B) {
	bDoAppend(1000000, b)
}

func BenchmarkDoIdx_10000000(b *testing.B) {
	bDoIdx(10000000, b)
}

func BenchmarkDoAppend_10000000(b *testing.B) {
	bDoAppend(10000000, b)
}
