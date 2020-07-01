# smparser
Matching wildcards string patterns based on state machine. Due to how expensive alternative pattern matching algorithms were, I wrote this lightweight finite state machine function to match wildcard strings.

## The algorithm
```
while not EOF and final state:
  if input char is equal state token:
    goto next token
  else:
    if state token is not wildcard:
      if no previous state wildcard:
        break loop and return false
      current state moves to previous wildcard
    else:
      next wildcard state is current token

    if input char is equal peek(next token):
      goto previous input char
      goto next token
    if peek(next token) is wildcard:
      goto next token

    goto next cursor
if reached final state, return true
```


For example, to match any string with the pattern ba*ab, where the inputs to each state will be: 0[b] 1[a] 2[*] 3[a] 4[b]. Given an input _barrack_ which goes from states: 0[b] -> 1[a] -> 2[rr] -> 3[a]. There is a backtrack to state 2 at input c and since it remains in this state until end of file, there's no match. Another input _barrackab_ would go from states: 0[b] -> 1[a] -> 2[rr] -> 3[a] -> 2[ck] -> 3[a] -> 4[b]. The input gets to the final state, hence it's a match.


## Testing
Benchmark test compares `smparser.MatchString` method with `regexp.MatchString`. regexp takes about 1.886s to compute while smparser takes about 0.459s.
