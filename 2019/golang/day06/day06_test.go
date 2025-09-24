package day06_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2019/golang/day06"
)

func TestPart1(t *testing.T) {
	d := day06.Parse("inputs/day06/sample-input.txt")
	got := d.Part1()
	want := 42
	if got != want {
		t.Errorf("Part1() = %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	d := day06.Parse("inputs/day06/sample-input2.txt")
	got := d.Part2()
	want := 4
	if got != want {
		t.Errorf("Part2() = %d; want %d", got, want)
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
		day06.Solve("inputs/day06/input.txt")
	}
}
