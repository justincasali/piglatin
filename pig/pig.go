// Package pig provides functions for converting text to pig latin
package pig

import (
	"strings"
	"unicode"
)

const (
	vowels  = "aeiou"
	vSuffix = "way"
	cSuffix = "ay"
)

func convert(input string) string {

	v := strings.IndexAny(input, vowels)

	// starts with vowel
	if v == 0 {
		return input + vSuffix
	}

	// starts with consontant
	if v > 0 {
		return input[v:] + input[:v] + cSuffix
	}

	return input

}

func split(input string) (output []string) {

	// start of chain
	start := 0

	// letter & other rune chains
	letter, other := false, false

	// loop through runes in input string
	for index, value := range input {

		// check if letter
		if unicode.IsLetter(value) {

			// end rune chain
			if other {
				output = append(output, input[start:index])
				start = index
			}

			// set rune chain
			letter, other = true, false

		} else {

			// end rune chain
			if letter {
				output = append(output, input[start:index])
				start = index
			}

			// set rune chain
			letter, other = false, true

		}

	}

	return append(output, input[start:])

}

// Latin converts text to pig latin
func Latin(input string) string {
	lower := strings.ToLower(input)

	var b strings.Builder
	b.Grow(len(lower) + len(lower)/4)

	for _, w := range split(lower) {
		b.WriteString(convert(w))
	}

	return b.String()
}
