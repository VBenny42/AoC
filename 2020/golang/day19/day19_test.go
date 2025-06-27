package day19_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day19"
)

func TestPart1(t *testing.T) {
	day := day19.Parse("inputs/day19/sample-input.txt")
	if got := day.Part1(); got != 2 {
		t.Errorf("Part1() = %v, want %v", got, 2)
	}
}

func TestPart2(t *testing.T) {
	day := day19.Parse("inputs/day19/sample-input1.txt")
	if got := day.Part2(); got != 12 {
		t.Errorf("Part1() = %v, want %v", got, 2)
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
		day19.Solve("inputs/day19/input.txt")
	}
}
