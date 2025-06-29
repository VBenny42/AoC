package day20

import "github.com/VBenny42/AoC/2020/golang/utils"

func (t *tile) rotate() {
	newData := make([]string, len(t.data))
	for i := range newData {
		newData[i] = ""
	}
	for i := len(t.data) - 1; i >= 0; i-- {
		for j := 0; j < len(t.data[i]); j++ {
			newData[j] += string(t.data[i][j])
		}
	}
	t.data = newData
	t.updateEdges()
}

func (t *tile) flipHorizontally() {
	newData := make([]string, len(t.data))
	for i, row := range t.data {
		newData[i] = utils.Reverse(row)
	}
	t.data = newData
	t.updateEdges()
}

func (t *tile) updateEdges() {
	// top
	t.edges[0] = t.data[0]
	// bottom
	t.edges[2] = t.data[len(t.data)-1]
	// left and right
	left, right := "", ""
	for i := 0; i < len(t.data); i++ {
		left += string(t.data[i][0])
		right += string(t.data[i][len(t.data[i])-1])
	}
	t.edges[3] = left
	t.edges[1] = right
}
