package day03_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2019/golang/day03"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"inputs/day03/sample-input.txt", 6},
		{"inputs/day03/sample-input2.txt", 159},
		{"inputs/day03/sample-input3.txt", 135},
	}

	for _, test := range tests {
		d := day03.Parse(test.filename)
		result, _ := d.Part1And2()
		if result != test.expected {
			t.Errorf("Part1(%s) = %d; want %d", test.filename, result, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		{"inputs/day03/sample-input.txt", 30},
		{"inputs/day03/sample-input2.txt", 610},
		{"inputs/day03/sample-input3.txt", 410},
	}

	for _, test := range tests {
		d := day03.Parse(test.filename)
		_, result := d.Part1And2()
		if result != test.expected {
			t.Errorf("Part2(%s) = %d; want %d", test.filename, result, test.expected)
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
		day03.Solve("inputs/day03/input.txt")
	}
}
