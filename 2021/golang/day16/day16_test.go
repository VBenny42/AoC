package day16_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day16"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		filepath string
		expected int
	}{
		{"inputs/day16/sample-input5.txt", 14},
		{"inputs/day16/sample-input1.txt", 16},
		{"inputs/day16/sample-input2.txt", 12},
		{"inputs/day16/sample-input3.txt", 23},
		{"inputs/day16/sample-input4.txt", 31},
	}

	for _, tc := range testCases {
		d := day16.Parse(tc.filepath)
		if got, want := d.Part1(), tc.expected; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		filepath string
		expected int
	}{
		{"inputs/day16/sample-input-part2-1.txt", 3},
		{"inputs/day16/sample-input-part2-2.txt", 54},
		{"inputs/day16/sample-input-part2-3.txt", 7},
		{"inputs/day16/sample-input-part2-4.txt", 9},
		{"inputs/day16/sample-input-part2-5.txt", 1},
		{"inputs/day16/sample-input-part2-6.txt", 0},
		{"inputs/day16/sample-input-part2-7.txt", 0},
		{"inputs/day16/sample-input-part2-8.txt", 1},
	}

	for _, tc := range testCases {
		d := day16.Parse(tc.filepath)
		if got, want := d.Part2(), tc.expected; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
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
		day16.Solve("inputs/day16/input.txt")
	}
}
