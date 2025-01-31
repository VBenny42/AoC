package day19

type interval struct {
	min, max int
}

type partRange struct {
	ranges [4]interval
}

func (pr partRange) split(c *condition) (pass, fail partRange, okPass, okFail bool) {
	pass = pr
	fail = pr

	current := pr.ranges[c.idx]

	if c.op == gt {
		passMin := max(current.min, c.value+1)
		passMax := current.max
		pass.ranges[c.idx] = interval{passMin, passMax}

		failMin := current.min
		failMax := min(current.max, c.value)
		fail.ranges[c.idx] = interval{failMin, failMax}
	} else {
		passMin := current.min
		passMax := min(current.max, c.value-1)
		pass.ranges[c.idx] = interval{passMin, passMax}

		failMin := max(current.min, c.value)
		failMax := current.max
		fail.ranges[c.idx] = interval{failMin, failMax}
	}

	okPass = pass.ranges[c.idx].min <= pass.ranges[c.idx].max
	okFail = fail.ranges[c.idx].min <= fail.ranges[c.idx].max

	return
}

func (d *day19) Part2() (combinations int) {
	type node struct {
		workflowName string
		pr           partRange
	}

	queue := []node{
		{
			workflowName: "in",
			pr: partRange{
				ranges: [4]interval{
					{1, 4000}, // x
					{1, 4000}, // m
					{1, 4000}, // a
					{1, 4000}, // s
				},
			},
		},
	}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		currWorkflow := entry.workflowName
		currRange := entry.pr

		if currWorkflow == "A" {
			product := 1
			for _, interval := range currRange.ranges {
				product *= (interval.max - interval.min + 1)
			}
			combinations += product
			continue
		}
		if currWorkflow == "R" {
			continue
		}

		wf, exists := d.workflows[currWorkflow]
		if !exists {
			continue
		}

		for _, rule := range wf.rules {
			if rule.cond == nil {
				queue = append(queue, node{rule.target, currRange})
				break
			}

			pass, fail, okPass, okFail := currRange.split(rule.cond)

			if okPass {
				queue = append(queue, node{rule.target, pass})
			}

			if okFail {
				currRange = fail
			} else {
				break
			}
		}
	}

	return
}
