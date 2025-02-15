package day20_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2021/golang/day20"
)

func TestPart1(t *testing.T) {
	d := day20.Parse("inputs/day20/sample-input.txt")

	want := 35
	if got := d.Part1(); got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
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
		day20.Solve("inputs/day20/input.txt")
	}
}
