package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
)

type packet struct {
	version    int
	typeID     int
	value      int
	subpackets []packet
}

type day16 struct {
	packet packet
}

func parseLiteralPacket(encodedPacket string, index *int) (p packet) {
	var reachedEnd bool
	var builder strings.Builder

	for !reachedEnd {
		if encodedPacket[*index] == '0' {
			reachedEnd = true
		}
		builder.WriteString(encodedPacket[*index+1 : *index+5])
		*index += 5
	}
	value, err := strconv.ParseInt(builder.String(), 2, 64)
	if err != nil {
		panic(fmt.Errorf("Error parsing value: %v", err))
	}

	p.value = int(value)
	return
}

func parseEncodedPacket(encodedPacket string, index *int) (p packet) {
	var lenId, length int

	if encodedPacket[*index] == '0' {
		lenId = 0
		*index++
		length64, err := strconv.ParseInt(encodedPacket[*index:*index+15], 2, 64)
		if err != nil {
			panic(fmt.Errorf("Error parsing length: %v", err))
		}
		length = int(length64)
		*index += 15
	} else {
		lenId = 1
		*index++
		length64, err := strconv.ParseInt(encodedPacket[*index:*index+11], 2, 64)
		if err != nil {
			panic(fmt.Errorf("Error parsing length: %v", err))
		}
		length = int(length64)
		*index += 11
	}

	var subpackets []packet

	if lenId == 0 {
		packetEnd := *index + length
		for *index < packetEnd {
			subpackets = append(subpackets, parsePacket(encodedPacket, index))
		}
	} else {
		for range length {
			subpackets = append(subpackets, parsePacket(encodedPacket, index))
		}
	}

	p.subpackets = subpackets
	return
}

func parsePacket(encodedPacket string, index *int) (p packet) {
	version, err := strconv.ParseInt(encodedPacket[*index:*index+3], 2, 64)
	if err != nil {
		panic(fmt.Errorf("Error parsing version: %v", err))
	}
	*index += 3

	typeID, err := strconv.ParseInt(encodedPacket[*index:*index+3], 2, 64)
	if err != nil {
		panic(fmt.Errorf("Error parsing typeID: %v", err))
	}
	*index += 3

	if typeID == 4 {
		p = parseLiteralPacket(encodedPacket, index)
	} else {
		p = parseEncodedPacket(encodedPacket, index)
	}

	p.version = int(version)
	p.typeID = int(typeID)

	return
}

const (
	sum = iota
	product
	minimum
	maximum
	value
	greaterThan
	lesserThan
	equalTo
)

func (p packet) evaluateExpression() int {
	switch p.typeID {
	case sum:
		sum := 0
		for _, subpacket := range p.subpackets {
			sum += subpacket.evaluateExpression()
		}
		return sum

	case product:
		product := 1
		for _, subpacket := range p.subpackets {
			product *= subpacket.evaluateExpression()
		}
		return product

	case minimum:
		minimum := p.subpackets[0].evaluateExpression()
		for _, subpacket := range p.subpackets[1:] {
			minimum = min(minimum, subpacket.evaluateExpression())
		}
		return minimum

	case maximum:
		maximum := p.subpackets[0].evaluateExpression()
		for _, subpacket := range p.subpackets[1:] {
			maximum = max(maximum, subpacket.evaluateExpression())
		}
		return maximum

	case value:
		return p.value

	case greaterThan:
		if p.subpackets[0].evaluateExpression() > p.subpackets[1].evaluateExpression() {
			return 1
		}
		return 0

	case lesserThan:
		if p.subpackets[0].evaluateExpression() < p.subpackets[1].evaluateExpression() {
			return 1
		}
		return 0

	case equalTo:
		if p.subpackets[0].evaluateExpression() == p.subpackets[1].evaluateExpression() {
			return 1
		}
		return 0

	default:
		// Not possible anyhow
		fmt.Println("ERROR: unknown typeID:", p.typeID)
		return -1
	}
}

func (p packet) sumVersions() int {
	sum := p.version
	for _, subpacket := range p.subpackets {
		sum += subpacket.sumVersions()
	}
	return sum
}

func (d *day16) Part1() int {
	return d.packet.sumVersions()
}

func (d *day16) Part2() int {
	return d.packet.evaluateExpression()
}

func Parse(filename string) *day16 {
	data := utils.ReadTrimmed(filename)

	var builder strings.Builder
	for _, c := range data {
		num, err := strconv.ParseInt(string(c), 16, 64)
		if err != nil {
			panic(err)
		}
		builder.WriteString(fmt.Sprintf("%04b", num))
	}

	var index int

	p := parsePacket(builder.String(), &index)

	return &day16{packet: p}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: sum of version numbers in all packets:", day.Part1())
	fmt.Println("ANSWER2: value of evaluated expression:", day.Part2())
}
