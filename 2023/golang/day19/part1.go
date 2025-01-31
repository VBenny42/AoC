package day19

func (c *condition) evaluate(p part) bool {
	switch c.op {
	case gt:
		return p[c.idx] > c.value
	case lt:
		return p[c.idx] < c.value
	}

	return false
}

func (d *day19) Part1() (sum int) {
	for _, p := range d.parts {
		currWorkflow := "in"

		for {
			if currWorkflow == "A" {
				for _, v := range p {
					sum += v
				}
				break
			}
			if currWorkflow == "R" {
				break
			}

			for _, r := range d.workflows[currWorkflow].rules {
				if r.cond != nil && r.cond.evaluate(p) {
					currWorkflow = r.target
					break
				}
				currWorkflow = r.target
			}
		}
	}

	return
}
