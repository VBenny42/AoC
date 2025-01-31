package day20_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day20"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day20.Parse("inputs/day20/sample-input.txt")

	assert.Equal(t, 11687500, d.Part1())
}

// // No part 2 sample input for this day.
// func ExampleSolve() {
// 	day20.Solve("inputs/day20/sample-input.txt")
// 	// Output:
// 	// ANSWER1: product of low pulses and high pulses: 11687500
// }

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
		day20.Solve("inputs/day20/input.txt")
	}
}
