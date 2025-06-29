package day14_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day14"
)

func TestPart1(t *testing.T) {
	expected := 165
	result := day14.Parse("inputs/day14/sample-input.txt").Part1()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 208
	result := day14.Parse("inputs/day14/sample-input1.txt").Part2()

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
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
		day14.Solve("inputs/day14/input.txt")
	}
}
