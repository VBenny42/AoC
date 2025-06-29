package day22

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
	"github.com/VBenny42/AoC/2020/golang/utils/set"
)

type day22 struct {
	player1 []int
	player2 []int
}

func (d *day22) playRound() bool {
	if len(d.player1) == 0 || len(d.player2) == 0 {
		return true // No cards left for either player
	}

	card1 := d.player1[0]
	card2 := d.player2[0]

	d.player1 = d.player1[1:]
	d.player2 = d.player2[1:]

	if card1 > card2 {
		d.player1 = append(d.player1, card1, card2)
		return false
	} else {
		d.player2 = append(d.player2, card2, card1)
		return false
	}
}

func (d *day22) playRecursiveGame(player1, player2 []int) (int, []int) {
	seen := set.NewSet[uint64]()

	for len(player1) > 0 && len(player2) > 0 {
		state := d.gameStateHash(player1, player2)
		if seen.Contains(state) {
			return 1, player1
		}
		seen.Add(state)

		card1 := player1[0]
		card2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]

		var player1Wins bool

		if len(player1) >= card1 && len(player2) >= card2 {
			subPlayer1 := make([]int, card1)
			subPlayer2 := make([]int, card2)
			copy(subPlayer1, player1[:card1])
			copy(subPlayer2, player2[:card2])

			winner, _ := d.playRecursiveGame(subPlayer1, subPlayer2)
			player1Wins = (winner == 1)
		} else {
			player1Wins = (card1 > card2)
		}

		if player1Wins {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}
	}

	if len(player1) > 0 {
		return 1, player1
	}
	return 2, player2
}

func (d *day22) gameStateHash(player1, player2 []int) uint64 {
	const fnvOffsetBasis = 0xcbf29ce484222325
	const fnvPrime = 0x100000001b3

	hash := uint64(fnvOffsetBasis)

	// Hash player1 cards
	for _, card := range player1 {
		hash ^= uint64(card)
		hash *= fnvPrime
	}

	// Add separator
	hash ^= 0xFFFF
	hash *= fnvPrime

	// Hash player2 cards
	for _, card := range player2 {
		hash ^= uint64(card)
		hash *= fnvPrime
	}

	return hash
}

func (d *day22) Part1() (score int) {
	for !d.playRound() {
	}

	var winningDeck []int

	winningDeck = d.player1
	if len(d.player2) > 0 {
		winningDeck = d.player2
	}

	for i, card := range winningDeck {
		score += card * (len(winningDeck) - i)
	}

	return
}

func (d *day22) Part2() (score int) {
	_, winningDeck := d.playRecursiveGame(d.player1, d.player2)

	// Calculate score for winning player
	for i, card := range winningDeck {
		score += card * (len(winningDeck) - i)
	}

	return
}

func Parse(filename string) *day22 {
	lines := utils.ReadLines(filename)
	var player1, player2 []int
	player1Done := false

	for _, line := range lines {
		if line == "" {
			player1Done = true
			continue
		}
		if strings.HasPrefix(line, "Player") {
			continue // Skip player header
		}
		if player1Done {
			player2 = append(player2, utils.Atoi(line))
		} else {
			player1 = append(player1, utils.Atoi(line))
		}
	}

	return &day22{player1: player1, player2: player2}
}

func Solve(filename string) {
	part1 := Parse(filename)
	part2 := day22{player1: part1.player1, player2: part1.player2}

	fmt.Println("ANSWER1: winning player's score:", part1.Part1())
	fmt.Println("ANSWER2: winning player's score for recursive game:", part2.Part2())
}
