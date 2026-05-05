package main

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/justincasali/piglatin/pig"
)

//go:embed data/test_input.txt
var testInput string

//go:embed data/test_output.txt
var testOutput string

func TestFunc(t *testing.T) {

	pigLatin := pig.Latin(testInput)
	if pigLatin != testOutput {
		t.Fail()
	}

}

func Example() {

	pigLatin := pig.Latin("Hello")
	fmt.Println(pigLatin)
	// Output: ellohay

}

func FuzzLatin(f *testing.F) {
	f.Add("")
	f.Add("Hello")
	f.Add("The quick brown fox jumps over the lazy dog")
	f.Add("123 !@# αβγ")
	f.Add(testInput)

	f.Fuzz(func(t *testing.T, s string) {
		out := pig.Latin(s)
		if !utf8.ValidString(out) {
			t.Fatalf("pig.Latin produced invalid UTF-8 for input %q: %q", s, out)
		}
	})
}

var (
	asciiInput   = strings.Repeat("The quick brown fox jumps over the lazy dog. ", 32)
	unicodeInput = strings.Repeat("The quick brown αβγ jumps over the lazy δog. ", 32)
)

func BenchmarkLatinASCII(b *testing.B) {
	for b.Loop() {
		_ = pig.Latin(asciiInput)
	}
}

func BenchmarkLatinUnicode(b *testing.B) {
	for b.Loop() {
		_ = pig.Latin(unicodeInput)
	}
}
