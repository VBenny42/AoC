package day25

import (
	"fmt"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type day25 struct {
	doorPublicKey int
	cardPublicKey int
}

const prime = 20201227

func getLoopSize(publicKey int) (loopSize int) {
	var (
		value         = 1
		subjectNumber = 7
	)

	for value != publicKey {
		value = (value * subjectNumber) % prime
		loopSize++
	}

	return
}

func (d *day25) Part1() int {
	// I know input is valid, so no need to actually check handshake validity.
	var (
		// doorLoopSize = getLoopSize(d.doorPublicKey)
		cardLoopSize = getLoopSize(d.cardPublicKey)

		doorEncryptionKey = 1
		// cardEncryptionKey = 1
	)

	for range cardLoopSize {
		doorEncryptionKey = (doorEncryptionKey * d.doorPublicKey) % prime
	}
	// for range doorLoopSize {
	// 	cardEncryptionKey = (cardEncryptionKey * d.cardPublicKey) % prime
	// }
	// if doorEncryptionKey != cardEncryptionKey {
	// 	panic("Encryption keys do not match")
	// }

	return doorEncryptionKey
}

func Parse(filename string) *day25 {
	lines := utils.ReadLines(filename)
	if len(lines) < 2 {
		panic("Not enough lines in input file")
	}

	return &day25{
		doorPublicKey: utils.Atoi(lines[0]),
		cardPublicKey: utils.Atoi(lines[1]),
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER: handshake encryption key:", day.Part1())
}
