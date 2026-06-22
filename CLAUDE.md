# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the CLI (no args uses a default sentence)
go run . "your text here"

# Run tests
go test ./...

# Run a single test
go test -run TestFunc .
go test -run Example .

# Run the fuzz test
go test -fuzz FuzzLatin -fuzztime 30s .

# Run benchmarks
go test -bench . .
```

## Architecture

This is a Go CLI and library for converting text to Pig Latin.

- **`pig/pig.go`** — the core library (`pig.Latin(input string) string`). This is the only logic worth touching.
- **`piglatin.go`** — thin `main` package; reads a CLI arg and calls `pig.Latin`.
- **`piglatin_test.go`** — tests, fuzz, and benchmarks live in `package main` using `//go:embed` to load `data/test_input.txt` / `data/test_output.txt` as golden files.

### `pig.Latin` algorithm

1. Lowercases the entire input up front (once, for stable Unicode handling).
2. Scans runes, splitting into maximal letter/non-letter chains.
3. For each letter chain: vowel-start → append `"way"`; consonant-start → rotate leading consonants to end, append `"ay"`; no vowels → pass through unchanged.
4. Non-letter chains (punctuation, spaces, digits) are copied through unchanged.

The golden files in `data/` are the source of truth for expected output — update both if the algorithm changes.
