package day16_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/day16"
)

func TestPart1(t *testing.T) {
	d := day16.Parse("inputs/day16/sample-input.txt")

	if d.Part1() != 71 {
		t.Errorf("expected 71, got %d", d.Part1())
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
