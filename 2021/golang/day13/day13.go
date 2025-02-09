package day13

import (
	"fmt"
	"image"
	"regexp"
	"strings"

	"github.com/VBenny42/AoC/2021/golang/utils"
	"github.com/VBenny42/AoC/2021/golang/utils/set"
)

type fold struct {
	foldOnX bool
	axis    int
}

type day13 struct {
	points set.Set[image.Point]
	folds  []fold
}

func (d *day13) foldPaper(index int) int {
	f := d.folds[index]

	nextSet := set.NewSet[image.Point]()

	if f.foldOnX {
		for c := range d.points {
			if c.X > f.axis {
				newPoint := image.Pt(f.axis-(c.X-f.axis), c.Y)
				nextSet.Add(newPoint)
			} else {
				nextSet.Add(c)
			}
		}
	} else {
		for c := range d.points {
			if c.Y > f.axis {
				newPoint := image.Pt(c.X, f.axis-(c.Y-f.axis))
				nextSet.Add(newPoint)
			} else {
				nextSet.Add(c)
			}
		}
	}

	d.points = nextSet

	return len(d.points)
}

func (d *day13) Part1() int {
	return d.foldPaper(0)
}

var imageFilename = "day13-letters.png"

func (d *day13) Part2() error {
	for i := 1; i < len(d.folds); i++ {
		d.foldPaper(i)
	}

	var maxX, maxY int
	for c := range d.points {
		maxX = max(maxX, c.X)
		maxY = max(maxY, c.Y)
	}

	g := grid(utils.NewGrid[rune](maxX+1, maxY+1))

	for c := range d.points {
		g[c.Y][c.X] = '#'
	}

	return g.writeImage(imageFilename)
}

func Parse(filename string) *day13 {
	var (
		data   = strings.SplitN(utils.ReadTrimmed(filename), "\n\n", 2)
		points = set.NewSet[image.Point]()
		folds  []fold
	)

	for _, line := range strings.Split(data[0], "\n") {
		point := strings.Split(line, ",")
		points.Add(image.Pt(utils.Atoi(point[0]), utils.Atoi(point[1])))
	}

	pattern := regexp.MustCompile(`(x|y)=(\d+)`)
	for _, line := range strings.Split(data[1], "\n") {
		parts := pattern.FindStringSubmatch(line)
		var f fold
		if parts[1] == "x" {
			f.foldOnX = true
		}
		f.axis = utils.Atoi(parts[2])
		folds = append(folds, f)
	}

	return &day13{
		points: points,
		folds:  folds,
	}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: dots visible after first fold instruction:", day.Part1())
	if err := day.Part2(); err != nil {
		fmt.Println("Error writing image:", err)
		return
	}
	fmt.Println("ANSWER2: letters visible after all fold instructions: written to", imageFilename)
}
