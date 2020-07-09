package smparser

import (
	"regexp"
	"testing"
)

func TestSmparserMatch(t *testing.T) {
	pattern := []string{
		"non*lu*ed*",
		"*iss*si*pi",
		"ba*ab*ab*ba",
		"*****da*****ad",
		"*nf*ig**on*",
		"home",
	}
	input := []string{
		"nonplussed",
		"mississipi",
		"balderdash absolutely abad bubba",
		"ahmedabad",
		"configuration",
		"home",
	}
	if len(pattern) != len(input) {
		t.Errorf("len error")
		return
	}
	for i := 0; i < len(pattern); i++ {
		ok := MatchString(pattern[i], input[i])
		if !ok {
			t.Errorf("expected a match")
			return
		}
	}
}

func TestSmparserNoMatch(t *testing.T) {
	pattern := []string{
		"non*lu*e*ed",
		"*iss*s*ssi*pi",
		"ba*ab*a*ba*ba",
		"**ed*a*b**b*ad",
		"*of*ig**on*",
		"this-is-a-string-that-is-not-short*-intentionally",
	}
	input := []string{
		"nonplussed",
		"mississipi",
		"balderdash absolutely bad bubba",
		"ahmedabad",
		"configuration",
		"this-is-a-string",
	}
	if len(pattern) != len(input) {
		t.Errorf("len error")
		return
	}
	for i := 0; i < len(pattern); i++ {
		ok := MatchString(pattern[i], input[i])
		if ok {
			t.Errorf("did not expect a match")
			return
		}
	}
}

func BenchmarkSmparser(b *testing.B) {
	const inputString = "https://www.examples.com/tutorials/hello-world/result.json"
	const pattern = "https://www*.examples.com/*/*/result.jso*"
	var ok bool
	for n := 0; n < b.N; n++ {
		ok = MatchString(pattern, inputString)
	}
	if !ok {
		b.Errorf("expected ok")
	}
}

func BenchmarkRegexp(b *testing.B) {
	const inputString = "https://www.examples.com/tutorials/hello-world/result.json"
	const pattern = "https://www.*.examples.com/.*/.*/result.jso.*"
	var ok bool
	for n := 0; n < b.N; n++ {
		ok, _ = regexp.MatchString(pattern, inputString)
	}
	if !ok {
		b.Errorf("expected ok")
	}
}
