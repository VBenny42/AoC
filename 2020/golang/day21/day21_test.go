package day21_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day21"
)

func ExampleSolve() {
	day21.Solve("inputs/day21/sample-input.txt")
	// Output:
	// ANSWER1: number of ingredients that contain no allergens: 5
	// ANSWER2: canonical dangerous ingredient list: mxmxvkd,sqjhc,fvjkl
}

func BenchmarkSolve(b *testing.B) {
	// b.ReportAllocs()

	oldStdout := os.Stdout
	null, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	os.Stdout = null

	defer func() {
		os.Stdout = oldStdout
		null.Close()
	}()

	b.ResetTimer()

	for range b.N {
		day21.Solve("inputs/day21/input.txt")
	}
}
