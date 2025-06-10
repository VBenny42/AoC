package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2020/golang/utils"
)

type passport map[string]string

type day04 struct {
	passports []passport
}

func (d *day04) Part1And2() (validCount, validValuesCount int) {
	var (
		requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
		validEyes      = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	)

	for _, p := range d.passports {
		hasAllRequired := true

		for _, field := range requiredFields {
			if p[field] == "" {
				hasAllRequired = false
				break
			}
		}

		if !hasAllRequired {
			continue
		}

		validCount++

		valid := true

		byr, err := strconv.Atoi(p["byr"])
		if err != nil || byr < 1920 || byr > 2002 {
			valid = false
		}

		if valid {
			iyr, err := strconv.Atoi(p["iyr"])
			if err != nil || iyr < 2010 || iyr > 2020 {
				valid = false
			}
		}

		if valid {
			eyr, err := strconv.Atoi(p["eyr"])
			if err != nil || eyr < 2020 || eyr > 2030 {
				valid = false
			}
		}

		if valid {
			hgt := p["hgt"]
			hgtValue, err := strconv.Atoi(hgt[:len(hgt)-2])
			switch {
			case strings.HasSuffix(hgt, "cm"):
				if err != nil || hgtValue < 150 || hgtValue > 193 {
					valid = false
				}
			case strings.HasSuffix(hgt, "in"):
				if err != nil || hgtValue < 59 || hgtValue > 76 {
					valid = false
				}
			default:
				valid = false
			}
		}

		if valid {
			hcl := p["hcl"]
			if len(hcl) != 7 || hcl[0] != '#' {
				valid = false
			} else {
				for _, c := range hcl[1:] {
					if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
						valid = false
						break
					}
				}
			}
		}

		if valid {
			ecl := p["ecl"]
			validEcl := false
			for _, validEye := range validEyes {
				if ecl == validEye {
					validEcl = true
					break
				}
			}
			if !validEcl {
				valid = false
			}
		}

		if valid {
			pid := p["pid"]
			if len(pid) != 9 {
				valid = false
			} else {
				for _, c := range pid {
					if c < '0' || c > '9' {
						valid = false
						break
					}
				}
			}
		}

		if valid {
			validValuesCount++
		}
	}
	return
}

func Parse(filename string) *day04 {
	var (
		passports []passport
		lines     = utils.ReadLines(filename)
		current   = make(passport)
	)

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				passports = append(passports, current)
				current = make(passport)
			}
		} else {
			fields := strings.Fields(line)
			for _, field := range fields {
				key, value, ok := strings.Cut(field, ":")
				if !ok {
					panic(fmt.Sprintf("invalid field format: %s", field))
				}
				current[key] = value
			}
		}
	}

	if len(current) > 0 {
		passports = append(passports, current)
	}

	return &day04{passports: passports}
}

func Solve(filename string) {
	day := Parse(filename)
	part1, part2 := day.Part1And2()

	fmt.Println("ANSWER1: number of valid passports:", part1)
	fmt.Println("ANSWER2: number of valid passports with valid values:", part2)
}
