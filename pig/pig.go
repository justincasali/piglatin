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

type chainKind uint8

const (
	// We treat the input as alternating chains of letters and non-letters.
	// This preserves punctuation/spacing exactly, while only transforming letters.
	noChain chainKind = iota
	letterChain
	otherChain
)

// Latin converts text to pig latin
func Latin(input string) string {
	// Lowercase once up-front so the rest of the algorithm operates on a single
	// consistent byte sequence. This keeps output stable for Unicode as well
	// (some lowercase mappings change UTF-8 byte length).
	lower := strings.ToLower(input)

	// Pre-grow with rough suffix headroom to reduce reallocations. Pig Latin adds
	// at most 3 bytes ("way") per transformed word.
	var b strings.Builder
	b.Grow(len(lower) + len(lower)/4)

	var (
		// Current chain kind and its starting byte offset in `lower`.
		kind  chainKind
		start int
	)

	flush := func(end int) {
		// Emit the current chain `lower[start:end]` into the builder.
		chain := lower[start:end]
		if kind != letterChain {
			// Non-letters are copied through unchanged.
			b.WriteString(chain)
			return
		}
		// For letter chains, find the first vowel and apply Pig Latin rules:
		// - starts with vowel: add "way"
		// - starts with consonant: rotate leading consonants to the end and add "ay"
		// - no vowels: leave as-is
		switch v := strings.IndexAny(chain, vowels); {
		case v == 0:
			b.WriteString(chain)
			b.WriteString(vSuffix)
		case v > 0:
			b.WriteString(chain[v:])
			b.WriteString(chain[:v])
			b.WriteString(cSuffix)
		default:
			b.WriteString(chain)
		}
	}

	for i, r := range lower {
		// Split the string into maximal letter/non-letter chains based on Unicode
		// category so behavior matches the prior implementation.
		k := otherChain
		if unicode.IsLetter(r) {
			k = letterChain
		}
		switch {
		case kind == noChain:
			// Start the first chain.
			kind, start = k, i
		case k != kind:
			// Chain boundary: flush previous chain and start the next.
			flush(i)
			kind, start = k, i
		}
	}
	if kind != noChain {
		// Flush the trailing chain (if any).
		flush(len(lower))
	}

	return b.String()
}
