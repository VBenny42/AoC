package solutions_test

import (
	"os"
	"testing"

	"github.com/VBenny42/AoC/2020/golang/solutions"
)

func BenchmarkAll(b *testing.B) {
	// b.ReportAllocs()

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

	b.ResetTimer()

	for range b.N {
		solutions.RunAll()
	}
}
