package smparser

// Matches returns true if the input string matches the pattern
func Matches(pattern string, str string) bool {
	EOF := len(str)
	finalState := len(pattern) - 1

	cursor := 0
	nextState := 0
	lastWild := -1

	for cursor != EOF && nextState <= finalState {
		if str[cursor] == pattern[nextState] {
			if nextState != finalState {
				nextState++
			}
		} else {
			if pattern[nextState] != '*' {
				if lastWild == -1 {
					break
				}
				nextState = lastWild
			} else {
				lastWild = nextState
			}
			if nextState+1 >= len(pattern) {
				// no lookahead
				break
			}
			lookahead := pattern[nextState+1]
			if str[cursor] == lookahead {
				cursor--
				nextState++
			}
			if lookahead == '*' {
				nextState++
			}
		}
		cursor++
	}
	matched := nextState == finalState

	return matched
}
