package day07_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2019/golang/day07"
)

func TestPart1(t *testing.T) {
	for _, tc := range []struct {
		filename string
		expected int
	}{
		{"inputs/day07/sample-input1.txt", 43210},
		{"inputs/day07/sample-input2.txt", 54321},
		{"inputs/day07/sample-input3.txt", 65210},
	} {
		t.Run(tc.filename, func(t *testing.T) {
			result := day07.Parse(tc.filename).Part1()
			if result != tc.expected {
				t.Errorf("Part1(%q) = %d; want %d", tc.filename, result, tc.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for _, tc := range []struct {
		filename string
		expected int
	}{
		{"inputs/day07/sample-input4.txt", 139629729},
		{"inputs/day07/sample-input5.txt", 18216},
	} {
		t.Run(tc.filename, func(t *testing.T) {
			result := day07.Parse(tc.filename).Part2()
			if result != tc.expected {
				t.Errorf("Part1(%q) = %d; want %d", tc.filename, result, tc.expected)
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
		day07.Solve("inputs/day07/input.txt")
	}
}
