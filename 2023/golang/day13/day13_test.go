package day13_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2023/golang/day13"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day13.Parse("inputs/day13/sample-input.txt")

	assert.Equal(t, 405, d.Part1())
}

func TestPart2(t *testing.T) {
	d := day13.Parse("inputs/day13/sample-input.txt")

	assert.Equal(t, 400, d.Part2())
}

func BenchmarkSolve(b *testing.B) {
	b.ReportAllocs()

	oldStdout := os.Stdout
	null, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}

	defer func() {
		os.Stdout = oldStdout
		null.Close()
	}()

	os.Stdout = null

	for range b.N {
		day13.Solve("inputs/day13/input.txt")
	}
}
