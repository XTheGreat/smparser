package smparser

// MatchString returns true if input string has an exact match to given pattern
func MatchString(pattern string, str string) bool {
	if pattern == str {
		return true
	}
	EOF := len(str)
	final := len(pattern)
	skip := -1
	for i := 0; i < final; i++ {
		if pattern[i] == '*' {
			skip = i
		}
	}
	if skip == -1 && pattern != str {
		return false
	}
	nstr := EOF - (final - skip) + 1
	if nstr < 0 {
		// input shorter than pattern
		return false
	}
	if skip != -1 && pattern[skip+1:final] != str[nstr:EOF] {
		return false
	}

	loopback := -1
	current := 0
	for cursor := 0; cursor < EOF && current != final; cursor++ {
		if str[cursor] == pattern[current] {
			current++
		} else if pattern[current] == '*' {
			loopback = current
			if current+1 < final && pattern[current+1] == str[cursor] ||
				current+1 < final && pattern[current+1] == '*' {
				current++
				cursor--
			}
		} else if loopback == -1 {
			break
		} else {
			current = loopback
			cursor--
		}
	}

	if current == skip && skip+1 == final {
		return pattern[current] == '*'
	}
	return current == final
}
