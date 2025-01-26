package day07

import (
	"strconv"
	"strings"
)

func (h hand) String() string {
	stringMap := map[card]string{
		ace:   "A",
		king:  "K",
		queen: "Q",
		joker: "*",
		jack:  "J",
		ten:   "T",
		nine:  "9",
		eight: "8",
		seven: "7",
		six:   "6",
		five:  "5",
		four:  "4",
		three: "3",
		two:   "2",
	}

	var hand strings.Builder
	for _, c := range h.cards {
		hand.WriteString(stringMap[c])
	}

	var set strings.Builder
	set.WriteString("{")
	for k, v := range h.set {
		set.WriteString(stringMap[k])
		set.WriteString(": ")
		set.WriteString(strconv.Itoa(v))
		set.WriteString(", ")
	}
	s := set.String()
	set.Reset()
	set.WriteString(s[:len(s)-2])
	set.WriteString("}")

	return hand.String() + " " + set.String() + " " + strconv.Itoa(h.bid) + " " + h.hType.String() + "\n"
}
