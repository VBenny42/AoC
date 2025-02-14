package day19

import (
	"fmt"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
	"github.com/VBenny42/AoC/2021/golang/utils/set"
)

type point struct {
	x, y, z int
}
type (
	beacon  = point
	scanner struct {
		id              int
		absolutePoint   point
		relativeBeacons []beacon
		absoluteBeacons []beacon
		distinctBeacons set.Set[beacon]
		rotations       [][]beacon
	}
)

type day19 struct {
	scanners []scanner
}

func vector(a, b beacon) beacon {
	return beacon{
		x: b.x - a.x,
		y: b.y - a.y,
		z: b.z - a.z,
	}
}

func (s *scanner) initializeRotations() {
	face1 := s.relativeBeacons
	var face2, face3, face4, face5, face6 []beacon
	for _, b := range face1 {
		face2 = append(face2, beacon{b.x, -b.y, -b.z})
		face3 = append(face3, beacon{b.x, -b.z, b.y})
		face4 = append(face4, beacon{-b.y, -b.z, b.x})
		face5 = append(face5, beacon{-b.x, -b.z, -b.y})
		face6 = append(face6, beacon{b.y, -b.z, -b.x})
	}
	rotations := [][]beacon{face1, face2, face3, face4, face5, face6}

	var actualRotations [][]beacon
	for _, r1 := range rotations {
		var r2, r3, r4 []beacon
		for _, b := range r1 {
			r2 = append(r2, beacon{-b.y, b.x, b.z})
			r3 = append(r3, beacon{-b.x, -b.y, b.z})
			r4 = append(r4, beacon{b.y, -b.x, b.z})
		}
		actualRotations = append(actualRotations, r1, r2, r3, r4)
	}

	s.rotations = actualRotations
}

func getAbsoluteBeacons(
	absolute, relative beacon,
	relativeBeacons []beacon,
) (beacons []beacon) {
	diff := vector(relative, absolute)

	for _, b := range relativeBeacons {
		beacons = append(beacons, beacon{
			x: diff.x + b.x,
			y: diff.y + b.y,
			z: diff.z + b.z,
		})
	}

	return
}

func findAbsoluteBeacons(unknown scanner, known []scanner) (scanner, bool) {
	for _, rotation := range unknown.rotations {
		for _, s := range known {
			for _, absoluteBeacon := range s.absoluteBeacons {
				for _, relativeBeacon := range rotation {
					unknownAbsoluteBeacons := getAbsoluteBeacons(absoluteBeacon, relativeBeacon, rotation)

					// Count matches against known scanner's absolute beacons
					matches := 0
					for _, newBeacon := range unknownAbsoluteBeacons {
						if s.distinctBeacons.Contains(newBeacon) {
							matches++
						}
					}

					if matches >= 12 {
						unknown.relativeBeacons = rotation
						unknown.absoluteBeacons = unknownAbsoluteBeacons
						unknown.distinctBeacons = set.NewSet(unknownAbsoluteBeacons...)
						unknown.absolutePoint = vector(relativeBeacon, absoluteBeacon)
						return unknown, true
					}
				}
			}
		}
	}

	return unknown, false
}

func (d *day19) Part1And2() (int, int) {
	known := []scanner{d.scanners[0]}
	known[0].absoluteBeacons = known[0].relativeBeacons
	known[0].distinctBeacons = set.NewSet(known[0].absoluteBeacons...)

	for i := range d.scanners {
		d.scanners[i].initializeRotations()
	}

	unknown := make([]scanner, len(d.scanners)-1)
	copy(unknown, d.scanners[1:])

	foundMatch := true
	for len(unknown) > 0 && foundMatch {
		foundMatch = false
		for i := 0; i < len(unknown); i++ {
			maybeKnown, ok := findAbsoluteBeacons(unknown[i], known)
			if ok {
				foundMatch = true
				known = append(known, maybeKnown)
				// Remove the matched scanner
				copy(unknown[i:], unknown[i+1:])
				unknown = unknown[:len(unknown)-1]
				break
			}
		}
	}

	allBeacons := set.NewSet[beacon]()
	for _, s := range known {
		for b := range s.distinctBeacons {
			allBeacons.Add(b)
		}
	}

	var furthest int
	for i, b1 := range known {
		for j, b2 := range known {
			if i == j {
				continue
			}

			manhattan := utils.Abs(b1.absolutePoint.x-b2.absolutePoint.x) +
				utils.Abs(b1.absolutePoint.y-b2.absolutePoint.y) +
				utils.Abs(b1.absolutePoint.z-b2.absolutePoint.z)
			furthest = max(furthest, manhattan)
		}
	}

	return len(allBeacons), furthest
}

func Parse(filename string) *day19 {
	var (
		data     = utils.ReadLines(filename)
		scanners []scanner
	)

	var (
		s  scanner
		id int
	)

	for i := 0; i < len(data); i++ {
		// Scanner line
		if strings.HasPrefix(data[i], "---") {
			s.id = id
			s.distinctBeacons = set.NewSet[beacon]()
			continue
		}
		// Block done
		if strings.TrimSpace(data[i]) == "" {
			id++
			scanners = append(scanners, s)
			s = scanner{}
			continue
		}

		split := strings.SplitN(data[i], ",", 3)
		if len(split) != 3 {
			panic(fmt.Sprintf("Invalid line: %s", data[i]))
		}

		b := beacon{
			x: utils.Atoi(split[0]),
			y: utils.Atoi(split[1]),
			z: utils.Atoi(split[2]),
		}

		s.relativeBeacons = append(s.relativeBeacons, b)
	}

	scanners = append(scanners, s)

	return &day19{scanners: scanners}
}

func Solve(filename string) {
	day := Parse(filename)

	part1, part2 := day.Part1And2()

	fmt.Println("ANSWER1: total beacons:", part1)
	fmt.Println("ANSWER2: largest manhattan distance between two scanners:", part2)
}
