package day09_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day09"
)

func TestPart1(t *testing.T) {
	day := day09.Parse("inputs/day09/sample-input.txt", 5)
	if got := day.Part1(); got != 127 {
		t.Errorf("Part1() = %d, want 127", got)
	}
}

func TestPart2(t *testing.T) {
	day := day09.Parse("inputs/day09/sample-input.txt", 5)
	day.Part1() // Ensure Part1 is called to set the part1 value
	if got := day.Part2(); got != 62 {
		t.Errorf("Part2() = %d, want 62", got)
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
		day09.Solve("inputs/day09/input.txt")
	}
}
