package day13_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day13"
)

func TestPart1(t *testing.T) {
	d := day13.Parse("inputs/day13/sample-input.txt")

	if got, want := d.Part1(), 17; got != want {
		t.Errorf("got %v, want %v", got, want)
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
		day13.Solve("inputs/day13/input.txt")
	}
}
