package day07_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day07"
)

func ExampleSolve() {
	day07.Solve("inputs/day07/sample-input.txt")
	// Output:
	// ANSWER1: fuel spent to get to optimal horizontal position: 37
	// ANSWER2: fuel spent to get to new optimal horizontal position: 168
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
		day07.Solve("inputs/day07/input.txt")
	}
}
