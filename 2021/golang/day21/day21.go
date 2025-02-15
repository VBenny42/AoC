package day21

import (
	"fmt"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

const (
	playerOne = iota
	playerTwo
)

type die struct {
	value int
}

type day21 struct {
	positions [2]int
	scores    [2]int
}

func (d *die) roll() (value int) {
	d.value++
	return d.value
}

func getRollAmount(die *die, rollAmount int) int {
	for range 3 {
		rollAmount += die.roll()
	}
	if rollAmount > 10 {
		rollAmount %= 10
	}
	if rollAmount == 0 {
		rollAmount = 10
	}
	return rollAmount
}

type key struct {
	positions, scores  [2]int
	rollsLeft          int
	currentIsPlayerOne bool
}

func diracPlay(sKey key, memo map[key][2]int64) (totalWins1, totalWins2 int64) {
	if v, ok := memo[sKey]; ok {
		return v[0], v[1]
	}

	index := playerTwo
	if sKey.currentIsPlayerOne {
		index = playerOne
	}

	scores := [2]int{sKey.scores[0], sKey.scores[1]}

	if sKey.rollsLeft == 0 {
		scores[index] += sKey.positions[index]

		if scores[index] >= 21 {
			if index == 0 {
				return 1, 0
			}
			return 0, 1
		}

		sKey.currentIsPlayerOne = !sKey.currentIsPlayerOne
		sKey.rollsLeft = 3

		index++
		index %= 2
	}

	for roll := range 3 {
		positions := [2]int{sKey.positions[0], sKey.positions[1]}
		positions[index] += roll + 1
		if positions[index] > 10 {
			positions[index] %= 10
		}
		w1, w2 := diracPlay(
			key{positions, scores, sKey.rollsLeft - 1, sKey.currentIsPlayerOne},
			memo,
		)
		totalWins1 += w1
		totalWins2 += w2
	}

	memo[sKey] = [2]int64{totalWins1, totalWins2}

	return totalWins1, totalWins2
}

func (d *day21) Part1() int {
	var rollAmount int
	die := new(die)

	for {
		rollAmount = getRollAmount(die, d.positions[playerOne])
		d.positions[playerOne] = rollAmount
		d.scores[playerOne] += rollAmount
		if d.scores[playerOne] >= 1000 {
			break
		}

		rollAmount = getRollAmount(die, d.positions[playerTwo])
		d.positions[playerTwo] = rollAmount
		d.scores[playerTwo] += rollAmount
		if d.scores[playerTwo] >= 1000 {
			break
		}
	}

	// At this point, either playerOne or playerTwo is the loser
	return min(d.scores[playerOne], d.scores[playerTwo]) * die.value
}

func (d *day21) Part2() int64 {
	return max(diracPlay(
		key{d.positions, d.scores, 3, true},
		make(map[key][2]int64),
	))
}

func Parse(filename string) *day21 {
	data := utils.ReadLines(filename)
	startingOne := utils.Atoi(data[0][len(data[0])-1:])
	startingTwo := utils.Atoi(data[1][len(data[1])-1:])

	return &day21{
		positions: [2]int{startingOne, startingTwo},
	}
}

func Solve(filename string) {
	fmt.Println("ANSWER1: losing play score times number of die rolls:", Parse(
		filename,
	).Part1())
	fmt.Println(
		"ANSWER2: number of universes where the most likely winner wins:",
		Parse(filename).Part2(),
	)
}
