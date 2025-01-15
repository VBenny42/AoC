package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"github.com/VBenny42/AoC/2022/golang/utils"
)

type robotType int

const (
	ore robotType = iota
	clay
	obsidian
	geode
)

var robots = []robotType{ore, clay, obsidian, geode}

type blueprint struct {
	costs      [4][3]int
	robotLimit [3]int
}

type state struct {
	timeRemaining int
	inventory     [3]int
	robots        [3]int
	geodes        int
}

type day19 struct {
	blueprints []blueprint
}

func startingState(timeRemaining int) *state {
	return &state{
		timeRemaining: timeRemaining,
		inventory:     [3]int{0, 0, 0},
		// Start off with 1 ore robot
		robots: [3]int{1, 0, 0},
		geodes: 0,
	}
}

func (s *state) canMakeRobot(blueprint blueprint, robot robotType) bool {
	for i := 0; i < 3; i++ {
		if s.inventory[i] < blueprint.costs[robot][i] {
			return false
		}
	}
	return true
}

func (s *state) makeAnotherRobot(blueprint blueprint, robot robotType) (newState state) {
	newState = state{
		timeRemaining: s.timeRemaining,
		geodes:        s.geodes,
	}
	copy(newState.inventory[:], s.inventory[:])
	copy(newState.robots[:], s.robots[:])

	for (!newState.canMakeRobot(blueprint, robot)) && (newState.timeRemaining > 1) {
		for i := 0; i < 3; i++ {
			newState.inventory[i] += newState.robots[i]
		}
		newState.timeRemaining--
	}

	newState.timeRemaining--

	costs := blueprint.costs[robot]
	for i := 0; i < 3; i++ {
		newState.inventory[i] = newState.inventory[i] - costs[i] + newState.robots[i]
	}

	if robot == geode {
		newState.geodes += newState.timeRemaining
	} else {
		newState.robots[robot]++
	}

	return
}

// Limit values were chosen by trial and error
// Just for optimization purposes, can use without
func (s *state) findMaxGeodes(blueprint blueprint, limits [4]int) int {
	if s.timeRemaining == 1 {
		return s.geodes
	}

	best := s.geodes
	for _, robot := range robots {
		if s.timeRemaining < limits[robot] ||
			(robot < geode && blueprint.robotLimit[robot] < s.robots[robot]) ||
			(robot == ore && s.robots[clay] > 1) ||
			(robot == obsidian && s.robots[clay] == 0) ||
			(robot == geode && s.robots[obsidian] == 0) {
			continue
		}

		nextState := s.makeAnotherRobot(blueprint, robot)
		if nextState.timeRemaining == 0 {
			continue
		}

		score := nextState.findMaxGeodes(blueprint, limits)
		best = max(best, score)
	}

	return best
}

func (d *day19) Part1() int {
	qualityLevelSum := 0

	var (
		wg      sync.WaitGroup
		levelCh = make(chan int)
	)
	wg.Add(len(d.blueprints))

	for i := range d.blueprints {
		go func(i int) {
			defer wg.Done()
			levelCh <- startingState(24).
				findMaxGeodes(d.blueprints[i], [4]int{18, 6, 3, 2}) * (i + 1)
		}(i)
	}

	go func() {
		wg.Wait()
		close(levelCh)
	}()

	for level := range levelCh {
		qualityLevelSum += level
	}

	return qualityLevelSum
}

func (d *day19) Part2() int {
	var (
		product = 1
		limit   = min(3, len(d.blueprints))
	)

	var (
		wg        sync.WaitGroup
		productCh = make(chan int)
	)
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go func(i int) {
			defer wg.Done()
			productCh <- startingState(32).
				findMaxGeodes(d.blueprints[i], [4]int{24, 10, 5, 2})
		}(i)
	}

	go func() {
		wg.Wait()
		close(productCh)
	}()

	for p := range productCh {
		product *= p
	}

	return product
}

var pattern = regexp.MustCompile(`(\d+)`)

func Parse(filename string) *day19 {
	data := utils.ReadLines(filename)

	blueprints := make([]blueprint, len(data))

	for i, line := range data {
		matches := pattern.FindAllString(line, -1)
		if len(matches) != 7 {
			panic("Invalid input")
		}

		blueprints[i].costs[ore][ore] = utils.Must(strconv.Atoi(matches[1]))
		blueprints[i].costs[clay][ore] = utils.Must(strconv.Atoi(matches[2]))
		blueprints[i].costs[obsidian][ore] = utils.Must(strconv.Atoi(matches[3]))
		blueprints[i].costs[obsidian][clay] = utils.Must(strconv.Atoi(matches[4]))
		blueprints[i].costs[geode][ore] = utils.Must(strconv.Atoi(matches[5]))
		blueprints[i].costs[geode][obsidian] = utils.Must(strconv.Atoi(matches[6]))

		blueprints[i].robotLimit[ore] = max(
			blueprints[i].costs[ore][ore],
			blueprints[i].costs[clay][ore],
			blueprints[i].costs[obsidian][ore],
			blueprints[i].costs[geode][ore],
		)
		blueprints[i].robotLimit[clay] = blueprints[i].costs[obsidian][clay]
		blueprints[i].robotLimit[obsidian] = blueprints[i].costs[geode][obsidian]
	}

	return &day19{blueprints}
}

func Solve(filename string) {
	day := Parse(filename)

	fmt.Println("ANSWER1: quality level of all blueprints:", day.Part1())
	fmt.Println("ANSWER2: product of the first three blueprints:", day.Part2())
}
