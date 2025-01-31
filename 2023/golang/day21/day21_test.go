package day21_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day21"
)

// func ExampleSolve() {
// 	day21.Solve("inputs/day21/sample-input.txt")
//   // Output:
//   // ANSWER1: <answer1>
// }

// func TestPart2(t *testing.T) {
// 	d := day21.Parse("inputs/day21/sample-input.txt")
//
// 	assert.Equal(t, 167004, d.Part2(500))
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
		day21.Solve("inputs/day21/input.txt")
	}
}
