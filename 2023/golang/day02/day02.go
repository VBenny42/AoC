package day02

import (
	"fmt"
	// "regexp"
	// "strconv"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type Sample struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Samples []Sample
}

type day02 struct {
	Games []Game
}

func (d *day02) Part1() (sum int) {
	limit := Sample{Red: 12, Green: 13, Blue: 14}

	for id, game := range d.Games {
		possible := true
		for _, sample := range game.Samples {
			if sample.Red > limit.Red ||
				sample.Green > limit.Green ||
				sample.Blue > limit.Blue {
				possible = false
				break
			}
		}

		if possible {
			sum += id + 1
		}
	}

	return
}

func (d *day02) Part2() (sum int) {
	for _, game := range d.Games {
		var (
			minRed   = game.Samples[0].Red
			minGreen = game.Samples[0].Green
			minBlue  = game.Samples[0].Blue
		)

		for _, sample := range game.Samples {
			minRed = max(minRed, sample.Red)
			minGreen = max(minGreen, sample.Green)
			minBlue = max(minBlue, sample.Blue)
		}

		sum += minRed * minGreen * minBlue
	}

	return
}

func Parse(filename string) *day02 {
	var (
		data  = utils.ReadLines(filename)
		games = make([]Game, len(data))
	)

	for i, line := range data {
		var (
			right         = strings.Split(line, ": ")[1]
			sampleStrings = strings.Split(right, "; ")
			samples       = make([]Sample, len(sampleStrings))
		)

		for j, sampleString := range sampleStrings {
			values := strings.Split(sampleString, ", ")

			for _, value := range values {
				numberAndColor := strings.Fields(value)
				if len(numberAndColor) != 2 {
					panic("Invalid number and color")
				}
				switch numberAndColor[1] {
				case "red":
					samples[j].Red = utils.Atoi(numberAndColor[0])
				case "green":
					samples[j].Green = utils.Atoi(numberAndColor[0])
				case "blue":
					samples[j].Blue = utils.Atoi(numberAndColor[0])
				}
			}
		}

		games[i] = Game{samples}
	}

	return &day02{games}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of IDs of possible games:", day.Part1())
	fmt.Println("ANSWER2: sum of power of sets:", day.Part2())
}
