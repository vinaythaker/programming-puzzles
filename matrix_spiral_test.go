package puzzles

import (
	"testing"
)

func TestBasicSpiral(t *testing.T) {
	w := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	req_answer := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	answer := spiralorder(w)
	for i := range req_answer {
		if answer[i] != req_answer[i] {
			t.Error("Answer does not match")
		} else {
			t.Log(answer)
		}
	}
}
