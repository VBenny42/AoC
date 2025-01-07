package day07

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type day07 struct {
	Directories map[string]int
}

func (d *day07) Part1() int {
	total := 0

	for _, size := range d.Directories {
		if size <= 100000 {
			total += size
		}
	}

	return total
}

func (d *day07) Part2() int {
	usedSpace := d.Directories["/"]
	unusedSpace := 70000000 - usedSpace
	minFree := usedSpace

	if unusedSpace < 30000000 {
		for _, size := range d.Directories {
			if size+unusedSpace >= 30000000 && size < minFree {
				minFree = size
			}
		}
	}

	return minFree
}

func Parse(filename string) *day07 {
	data := utils.ReadLines(filename)

	directories := make(map[string]int)
	var path []string
	var currentPath string
	var currentSize int

	for _, line := range data {
		fields := strings.Fields(line)
		if len(fields) == 3 { // cd commands
			if fields[2] == "/" {
				path = []string{fields[2]}
				currentPath = fields[2]
				directories[currentPath] = 0
			} else if fields[2] == ".." {
				currentSize = directories[currentPath]
				path = path[:len(path)-1]
				currentPath = strings.Join(path, "")
				directories[currentPath] += currentSize
			} else {
				path = append(path, fields[2]+"/")
				currentPath = strings.Join(path, "")
				directories[currentPath] = 0
			}
		}
		if len(fields) == 2 {
			if fields[0] == "dir" {
				dirPath := currentPath + fields[1] + "/"
				directories[dirPath] = 0
			} else if num, err := strconv.Atoi(fields[0]); err == nil {
				directories[currentPath] += num
			}
		}
	}

	for i := len(path) - 1; i > 0; i-- {
		currentSize = directories[currentPath]
		path = path[:i]
		currentPath = strings.Join(path, "")
		directories[currentPath] += currentSize
	}

	return &day07{directories}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: total size of directories <= 100000:", day.Part1())
	fmt.Println("ANSWER2: smallest directory to free:", day.Part2())
}
