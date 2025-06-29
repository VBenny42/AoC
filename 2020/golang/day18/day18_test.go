package day18_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day18"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"inputs/day18/sample-input1.txt", 71},
		{"inputs/day18/sample-input2.txt", 51},
		{"inputs/day18/sample-input3.txt", 26},
		{"inputs/day18/sample-input4.txt", 437},
		{"inputs/day18/sample-input5.txt", 12240},
		{"inputs/day18/sample-input6.txt", 13632},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			day := day18.Parse(test.filename)
			result := day.Part1()
			if result != test.expected {
				t.Errorf("Part1(%s) = %d; want %d", test.filename, result, test.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"inputs/day18/sample-input1.txt", 231},
		{"inputs/day18/sample-input2.txt", 51},
		{"inputs/day18/sample-input3.txt", 46},
		{"inputs/day18/sample-input4.txt", 1445},
		{"inputs/day18/sample-input5.txt", 669060},
		{"inputs/day18/sample-input6.txt", 23340},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			day := day18.Parse(test.filename)
			result := day.Part2()
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
		day18.Solve("inputs/day18/input.txt")
	}
}
