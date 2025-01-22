package day19

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2024/golang/embeds"
)

type day19 struct {
	towels  []string
	designs []string
	cache   map[string]int
}

func (d *day19) differentCombos(design string) int {
	if len(design) == 0 {
		return 1
	}

	if val, ok := d.cache[design]; ok {
		return val
	}

	count := 0
	for _, towel := range d.towels {
		if len(towel) <= len(design) && design[:len(towel)] == towel {
			combos := d.differentCombos(design[len(towel):])
			count += combos
			d.cache[design] = count
		}
	}
	return count
}

func (d *day19) Part1and2() (int, int) {
	possibleDesigns := 0
	allCombos := 0

	for _, design := range d.designs {
		combos := d.differentCombos(design)
		if combos > 0 {
			possibleDesigns++
		}
		allCombos += combos
	}

	return possibleDesigns, allCombos
}

func Parse(filename string) *day19 {
	file, err := embeds.Inputs.Open(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Scan()
	towels := strings.Split(s.Text(), ", ")
	s.Scan()

	designs := make([]string, 0)

	for s.Scan() {
		designs = append(designs, s.Text())
	}

	cache := make(map[string]int)

	return &day19{towels, designs, cache}
}

func Solve(filename string) {
	possibleDesigns, allCombos := Parse(filename).Part1and2()
	fmt.Println("ANSWER1: possibleDesigns:", possibleDesigns)
	fmt.Println("ANSWER2: allCombos:", allCombos)
}
