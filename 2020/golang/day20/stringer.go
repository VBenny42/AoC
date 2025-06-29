package day20

import (
	"fmt"
	"strconv"
	"strings"
)

func (t tile) String() string {
	return "Tile " + strconv.Itoa(t.id) + " Edges: " + fmt.Sprintf("%v", t.edges)
}

func (d *day20) String() string {
	var builder strings.Builder

	builder.WriteString("Day 20 Tiles:\n")
	for _, t := range d.tiles {
		builder.WriteString(t.String() + "\n")
	}

	return builder.String()
}
