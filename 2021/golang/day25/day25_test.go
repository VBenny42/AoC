package day25_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day25"
)

func TestPart1(t *testing.T) {
	d := day25.Parse("inputs/day25/sample-input.txt")

	if got, want := d.Part1(), 58; got != want {
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
		day25.Solve("inputs/day25/input.txt")
	}
}
