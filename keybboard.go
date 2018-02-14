package puzzles

import (
	"regexp"
	"strings"
)

func findWords(words []string) []string {

	output := make([]string, len(words))
	index := 0
	rex := regexp.MustCompile(`^[asdfghjkl]*$|^[qwertyuiop]*$|^[zxcvbnm]*$`)
	for _, i := range words {
		if rex.MatchString(strings.ToLower(i)) {
			output[index] = i
			index++
		}
	}
	return output[0:index]
}
