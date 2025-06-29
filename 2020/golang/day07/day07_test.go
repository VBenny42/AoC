package day07_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day07"
)

func ExampleSolve() {
	day07.Solve("inputs/day07/sample-input.txt")
	day07.Solve("inputs/day07/sample-input1.txt")
	// Output:
	// ANSWER1: bag colors that can eventually contain a shiny gold bag: 4
	// ANSWER2: number of bags inside a shiny gold bag: 32
	// ANSWER1: bag colors that can eventually contain a shiny gold bag: 0
	// ANSWER2: number of bags inside a shiny gold bag: 126
}

func BenchmarkSolve(b *testing.B) {
	b.ReportAllocs()

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
