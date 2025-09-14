package day04

import (
	"fmt"
	// "sync"

	"github.com/VBenny42/AoC/2019/golang/utils"
)

type day04 struct {
	start, end int
}

type password [6]int

// func makePassword(n int) password {
// 	var p password
// 	for i := 5; i >= 0; i-- {
// 		p[i] = n % 10
// 		n /= 10
// 	}
// 	return p
// }
//
// func (p *password) isNonDecreasing() bool {
// 	for i := 1; i < 6; i++ {
// 		if p[i] < p[i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func (p *password) isValid() (valid bool) {
// 	for i := 1; i < 6; i++ {
// 		if p[i] == p[i-1] {
// 			valid = true
// 		}
// 	}
//
// 	return
// }
//
// func (p *password) isValid2() (valid bool) {
// 	var count int
// 	for i := 1; i < 6; i++ {
// 		if p[i] == p[i-1] {
// 			count++
// 		} else {
// 			if count == 1 {
// 				valid = true
// 			}
// 			count = 0
// 		}
// 	}
// 	if count == 1 {
// 		valid = true
// 	}
//
// 	return
// }
//
// func (d *day04) Part1And2() (count1, count2 int) {
// 	const numWorkers = 100
// 	var (
// 		wg       sync.WaitGroup
// 		resultCh = make(chan [2]int, numWorkers)
// 		jobCh    = make(chan int, numWorkers)
// 	)
//
// 	// Start workers
// 	wg.Add(numWorkers)
// 	for i := 0; i < numWorkers; i++ {
// 		go func() {
// 			defer wg.Done()
// 			localCount1, localCount2 := 0, 0
// 			for n := range jobCh {
// 				p := makePassword(n)
// 				if !p.isNonDecreasing() {
// 					continue
// 				}
// 				if p.isValid() {
// 					localCount1++
// 					if p.isValid2() {
// 						localCount2++
// 					}
// 				}
// 			}
// 			resultCh <- [2]int{localCount1, localCount2}
// 		}()
// 	}
//
// 	// Send jobs
// 	go func() {
// 		for n := d.start; n <= d.end; n++ {
// 			jobCh <- n
// 		}
// 		close(jobCh)
// 	}()
//
// 	// Close result channel after all workers finish
// 	go func() {
// 		wg.Wait()
// 		close(resultCh)
// 	}()
//
// 	// Collect results
// 	for result := range resultCh {
// 		count1 += result[0]
// 		count2 += result[1]
// 	}
//
// 	return
// }

func (p *password) isValid() (valid bool) {
	for i := 1; i < 6; i++ {
		if p[i] < p[i-1] {
			return false
		}
		if p[i] == p[i-1] {
			valid = true
		}
	}
	return
}

func (p *password) isValid2() (valid bool) {
	var count int
	for i := 1; i < 6; i++ {
		if p[i] == p[i-1] {
			count++
		} else {
			if count == 1 {
				valid = true
			}
			count = 0
		}
	}
	if count == 1 {
		valid = true
	}
	return
}

func (day *day04) Part1And2() (count1, count2 int) {
	for a := 1; a <= 9; a++ {
		for b := a; b <= 9; b++ {
			for c := b; c <= 9; c++ {
				for d := c; d <= 9; d++ {
					for e := d; e <= 9; e++ {
						for f := e; f <= 9; f++ {
							num := a*100000 + b*10000 + c*1000 + d*100 + e*10 + f
							if num < day.start || num > day.end {
								continue
							}
							p := password{a, b, c, d, e, f}
							if p.isValid() {
								count1++
							}
							if p.isValid2() {
								count2++
							}
						}
					}
				}
			}
		}
	}
	return
}

func Parse(filename string) *day04 {
	var (
		day  day04
		line = utils.ReadTrimmed(filename)
	)

	fmt.Sscanf(line, "%d-%d", &day.start, &day.end)

	return &day
}

func Solve(filename string) {
	day := Parse(filename)

	count1, count2 := day.Part1And2()

	fmt.Println("ANSWER1: number of valid passwords:", count1)
	fmt.Println("ANSWER2: number of valid passwords with new double constraint:", count2)
}
