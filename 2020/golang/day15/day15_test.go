package day15_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day15"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"inputs/day15/sample-input1.txt", 436},
		{"inputs/day15/sample-input2.txt", 1},
		{"inputs/day15/sample-input3.txt", 10},
		{"inputs/day15/sample-input4.txt", 27},
		{"inputs/day15/sample-input5.txt", 78},
		{"inputs/day15/sample-input6.txt", 438},
		{"inputs/day15/sample-input7.txt", 1836},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			result := day15.Parse(test.filename).Part1()
			if result != test.expected {
				t.Errorf("Part1(%s) = %d; want %d", test.filename, result, test.expected)
			}
		})
	}
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
		day15.Solve("inputs/day15/input.txt")
	}
}
