package day25

import "strings"

func (n node) String() string {
	return string(n)
}

// Using for graphviz
func (g graph) String() string {
	var builder strings.Builder

	builder.WriteString("graph {\n")
	for e := range g.edges {
		builder.WriteString(e.n1.String() + " -- " + e.n2.String() + "\n")
	}
	builder.WriteString("}")

	return builder.String()
}
