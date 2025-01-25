package day02

import "fmt"

func (s Sample) String() string {
	return fmt.Sprintf("R: %d, G: %d, B: %d", s.Red, s.Green, s.Blue)
}

func (g Game) String() string {
	return fmt.Sprintf("Samples: %v\n", g.Samples)
}
