package day07

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
	"github.com/VBenny42/AoC/2020/golang/utils/set"
)

type colorIndex struct {
	colorToIndex map[string]int
	nextIndex    int
}

type day07 struct {
	bags map[int]map[int]int
}

const shinyGoldIndex = 0

func (ci *colorIndex) getIndex(color string) int {
	if index, exists := ci.colorToIndex[color]; exists {
		return index
	}

	index := ci.nextIndex
	ci.colorToIndex[color] = index
	ci.nextIndex++
	return index
}

func (d *day07) Part1() (count int) {
	uniqueBags := set.NewSet[int]()

	var canContainGold func(bag int) bool
	canContainGold = func(bag int) bool {
		if uniqueBags.Contains(bag) {
			return true
		}

		for color := range d.bags[bag] {
			if color == shinyGoldIndex || canContainGold(color) {
				uniqueBags.Add(bag)
				return true
			}
		}
		return false
	}

	for bagColor := range d.bags {
		if bagColor != shinyGoldIndex && canContainGold(bagColor) {
			uniqueBags.Add(bagColor)
		}
	}

	return len(uniqueBags)
}

func (d *day07) Part2() int {
	var countBags func(bag int) int
	countBags = func(bag int) int {
		total := 0
		for color, quantity := range d.bags[bag] {
			total += quantity + quantity*countBags(color)
		}
		return total
	}

	return countBags(shinyGoldIndex)
}

func Parse(filename string) *day07 {
	lines := utils.ReadLines(filename)
	bags := make(map[int]map[int]int, len(lines))
	colorIndex := colorIndex{
		colorToIndex: map[string]int{"shiny gold": shinyGoldIndex},
		nextIndex:    1,
	}

	var (
		parts    []string
		bagColor int

		contents       []string
		count          int
		color1, color2 string
	)

	for _, line := range lines {
		parts = strings.Split(line, " bags contain ")
		bagColor = colorIndex.getIndex(parts[0])
		bags[bagColor] = make(map[int]int)

		if parts[1] == "no other bags." {
			continue
		}

		contents = strings.Split(parts[1], ", ")
		for _, content := range contents {
			fmt.Sscanf(content, "%d %s %s bag", &count, &color1, &color2)
			bags[bagColor][colorIndex.getIndex(color1+" "+color2)] = count
		}
	}

	return &day07{bags: bags}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println(
		"ANSWER1: bag colors that can eventually contain a shiny gold bag:",
		day.Part1(),
	)
	fmt.Println(
		"ANSWER2: number of bags inside a shiny gold bag:",
		day.Part2(),
	)
}
