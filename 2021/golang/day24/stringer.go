package day24

import "fmt"

func (i instruction) String() string {
	opcodes := map[opcode]string{
		add: "add",
		mul: "mul",
		div: "div",
		mod: "mod",
		eql: "eql",
	}
	return fmt.Sprintf("%v %v %v", opcodes[i.opcode], i.a, i.b)
}
