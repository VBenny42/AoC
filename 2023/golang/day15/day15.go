package day15

import (
	"fmt"
	"slices"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type day15 struct {
	steps []string
}

func hash(s string) (hash int) {
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash %= 256
	}

	return
}

func (d *day15) Part1() (sum int) {
	for _, step := range d.steps {
		sum += hash(step)
	}

	return
}

type operation struct {
	label     string
	labelHash int
	length    *int
}

func (d *day15) Part2() (power int) {
	var (
		boxes      = make([][]string, 256)
		boxesMap   = make([]map[string]int, 256)
		operations = make([]operation, len(d.steps))
	)

	for i := range boxesMap {
		boxesMap[i] = make(map[string]int)
	}

	labelFunc := func(r rune) bool {
		return r == '=' || r == '-'
	}

	for i, step := range d.steps {
		fields := strings.FieldsFunc(step, labelFunc)
		operations[i] = operation{label: fields[0], labelHash: hash(fields[0])}
		if len(fields) == 2 {
			length := utils.Atoi(fields[1])
			operations[i].length = &length
		}
	}

	for _, op := range operations {
		// - label case, remove lens from box
		if op.length == nil {
			_, ok := boxesMap[op.labelHash][op.label]
			if !ok {
				continue
			}

			// remove lens from box
			delete(boxesMap[op.labelHash], op.label)
			index := slices.Index(boxes[op.labelHash], op.label)
			boxes[op.labelHash] = append(
				boxes[op.labelHash][:index],
				boxes[op.labelHash][index+1:]...,
			)
		} else {
			// = lens case, add lens to box
			_, ok := boxesMap[op.labelHash][op.label]

			if !ok {
				// no lens with label in box, add it
				boxes[op.labelHash] = append(boxes[op.labelHash], op.label)
				boxesMap[op.labelHash][op.label] = *op.length
			} else {
				// lens with label in box, update it
				boxesMap[op.labelHash][op.label] = *op.length
			}
		}
	}

	for boxIdx := range boxes {
		for slotIdx := range boxes[boxIdx] {
			power += (boxIdx + 1) *
				(slotIdx + 1) *
				boxesMap[boxIdx][boxes[boxIdx][slotIdx]]
		}
	}

	return
}

func Parse(filename string) *day15 {
	data := utils.ReadTrimmed(filename)

	return &day15{
		steps: strings.Split(data, ","),
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of hashes:", day.Part1())
	fmt.Println("ANSWER2: focusing power of lens configuration:", day.Part2())
}
