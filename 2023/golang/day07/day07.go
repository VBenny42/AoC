package day07

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/VBenny42/AoC/2023/golang/utils"
)

type (
	card     int
	handType int
)

type hand struct {
	cards [5]card
	set   map[card]int
	bid   int
	hType handType
}

const (
	joker card = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

var values = map[rune]card{
	'A': ace, 'K': king, 'Q': queen,
	'J': jack, 'T': ten, '9': nine,
	'8': eight, '7': seven, '6': six,
	'5': five, '4': four, '3': three, '2': two,
}

//go:generate stringer -type=handType
const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type day07 struct {
	hands []hand
}

func (d *day07) calculateWinnings() (winnings int) {
	slices.SortFunc(d.hands, func(a, b hand) int {
		if a.hType != b.hType {
			return cmp.Compare(a.hType, b.hType)
		}
		for i := range a.cards {
			if a.cards[i] != b.cards[i] {
				return cmp.Compare(a.cards[i], b.cards[i])
			}
		}
		return 0
	})

	for i, h := range d.hands {
		winnings += h.bid * (i + 1)
	}

	return
}

func (h *hand) makeBetterType() {
	jokerCount := h.set[joker]

	// Find the most frequent non-joker card
	bestCard := joker
	maxCount := 0
	for k, v := range h.set {
		if k != joker && v > maxCount {
			bestCard = k
			maxCount = v
		}
	}

	// Add joker count to most frequent card
	h.set[bestCard] += jokerCount
	delete(h.set, joker)

	// Recalculate hand type
	h.determineHandType()
}

func (d *day07) Part1() int {
	return d.calculateWinnings()
}

func (d *day07) Part2() int {
	for i := range d.hands {
		jackPresent := false
		for k := range d.hands[i].set {
			if k == jack {
				jackPresent = true
				break
			}
		}
		// no joker in hand, type cannot be changed
		if !jackPresent {
			continue
		}

		// replace all jacks with jokers
		for j := range d.hands[i].cards {
			if d.hands[i].cards[j] == jack {
				d.hands[i].cards[j] = joker
			}
		}

		jackCount := d.hands[i].set[jack]
		delete(d.hands[i].set, jack)
		d.hands[i].set[joker] = jackCount

		// hand is already the highest type,
		// no need to try and upgrade it
		if d.hands[i].hType == fiveOfAKind {
			continue
		}

		d.hands[i].makeBetterType()
	}

	return d.calculateWinnings()
}

func Parse(filename string) *day07 {
	data := utils.ReadLines(filename)

	hands := make([]hand, len(data))

	for i, line := range data {
		fields := strings.Fields(line)
		hands[i].set = make(map[card]int)
		for j, c := range fields[0] {
			c, ok := values[c]
			if !ok {
				panic("Invalid card value")
			}
			hands[i].cards[j] = c
			hands[i].set[c]++
		}
		hands[i].bid = utils.Atoi(fields[1])
		hands[i].determineHandType()
	}

	return &day07{hands}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: total winnings:", day.Part1())
	fmt.Println("ANSWER2: total winnings with jokers as wildcards:", day.Part2())
}
