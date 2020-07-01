package smparser

import (
	"regexp"
	"testing"
)

const inputString = "babar738egdjhjab"
const pattern = "ba*ab"
const TRIAL = 1000000

func BenchmarkSmparser(b *testing.B) {
	for n := 0; n < TRIAL; n++ {
		MatchString(pattern, inputString)
	}
}

func BenchmarkRegexp(b *testing.B) {
	for n := 0; n < TRIAL; n++ {
		regexp.MatchString(pattern, inputString)
	}
}
