package day05_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day05"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"inputs/day05/sample-input1.txt", 357},
		{"inputs/day05/sample-input2.txt", 567},
		{"inputs/day05/sample-input3.txt", 119},
		{"inputs/day05/sample-input4.txt", 820},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			day := day05.Parse(test.filename)
			result := day.Part1()
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
		day05.Solve("inputs/day05/input.txt")
	}
}
