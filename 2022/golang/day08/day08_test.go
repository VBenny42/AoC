package day08_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2022/golang/day08"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	d := day08.Parse("../../inputs/day08/sample-input.txt")

	value, _ := d.Part1and2()

	assert.Equal(t, 21, value)
}

func TestPart2(t *testing.T) {
	d := day08.Parse("../../inputs/day08/sample-input.txt")

	_, value := d.Part1and2()

	assert.Equal(t, 8, value)
}

func BenchmarkSolve(b *testing.B) {
	b.StopTimer()
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatalf("Failed to open %s: %s", os.DevNull, err)
	}
	defer devNull.Close()
	origStdout := os.Stdout
	os.Stdout = devNull
	b.StartTimer()

	for range b.N {
		day08.Solve("../../inputs/day08/sample-input.txt")
	}

	os.Stdout = origStdout
}
