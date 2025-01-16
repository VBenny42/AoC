package day20

import (
	"fmt"
	"strconv"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type day20 struct {
	file []*int
}

func (d *day20) mix(mixNum int) (mixed []*int, zeroIndex int) {
	mixed = make([]*int, len(d.file))
	copy(mixed, d.file)

	indexOf := func(arr []*int, val *int) int {
		for i, v := range arr {
			if val == v {
				return i
			}
		}
		return -1
	}

	var (
		lastIndex    = len(d.file) - 1
		zeroIndexPtr *int
	)

	for ; mixNum > 0; mixNum-- {
		for _, node := range d.file {
			oldIndex := indexOf(mixed, node)
			newIndex := (oldIndex + *node) % lastIndex

			if newIndex < 0 {
				newIndex += lastIndex
			}

			// 0 is a special case, shouldn't move
			if *node == 0 {
				zeroIndexPtr = node
				continue
			}

			// Remove the node at oldIndex
			copy(mixed[oldIndex:], mixed[oldIndex+1:])

			// Insert the node at newIndex
			copy(mixed[newIndex+1:], mixed[newIndex:])
			mixed[newIndex] = node
		}
	}

	zeroIndex = indexOf(mixed, zeroIndexPtr)

	return
}

func (d *day20) Part1() (sum int) {
	mixed, zeroIndex := d.mix(1)

	sum += *mixed[(zeroIndex+1000)%len(mixed)]
	sum += *mixed[(zeroIndex+2000)%len(mixed)]
	sum += *mixed[(zeroIndex+3000)%len(mixed)]

	return
}

func (d *day20) Part2() (sum int) {
	for _, node := range d.file {
		*node *= 811589153
	}

	mixed, zeroIndex := d.mix(10)

	sum += *mixed[(zeroIndex+1000)%len(mixed)]
	sum += *mixed[(zeroIndex+2000)%len(mixed)]
	sum += *mixed[(zeroIndex+3000)%len(mixed)]

	return
}

func Parse(filename string) *day20 {
	data := utils.ReadLines(filename)

	file := make([]*int, len(data))

	for i, line := range data {
		val := utils.Must(strconv.Atoi(line))
		file[i] = &val
	}

	return &day20{file}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of grove coordinates:", day.Part1())
	fmt.Println("ANSWER2: sum of grove coordinates after decrypting and mixing:",
		day.Part2())
}
