package day05_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2019/golang/day05"
)

func TestPart2(t *testing.T) {
	day := day05.Parse("inputs/day05/sample-input2.txt")
	for _, tc := range []struct {
		input    int
		expected int
	}{
		{input: 7, expected: 999},
		{input: 8, expected: 1000},
		{input: 9, expected: 1001},
	} {
		t.Run("", func(t *testing.T) {
			if got := day.Part2(tc.input); got != tc.expected {
				t.Errorf("Part2(%d) = %d; want %d", tc.input, got, tc.expected)
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
