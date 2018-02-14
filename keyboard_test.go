package puzzles

import (
	"testing"
)

func TestMixedRows(t *testing.T) {
	w := []string{"Hello", "Alaska", "Dad", "Peace"}
	answer := findWords(w)
	expected_answer := []string{"Alaska", "Dad"}
	for i := range expected_answer {
		if answer[i] != expected_answer[i] {
			t.Error("Answer does not match")
		} else {
			t.Log(answer)
		}
	}
}
