# smparser
Simpler lightweight function to match wildcard patterns faster.

## The algorithm
This function seeks to find the best way to return quickly. It starts by comparing the characters from last found wildcard in the pattern with the last tokens of the input and returns false if there's no match. Then it parses the input in a state machine and compares read input with state token. The function evaluates to true if final state/token is reached.

```
while not (EOF or out of token):
  if read input is equal token:
  	goto (next token)
  else if token is wildcard:
  	if next pattern token not out of range and is either wildcard or input char:
		goto (next token)
		goto (same input)
  else if no loopback:
  	return
	else:
		goto (loopback token)
		goto (same input)

evaluate match (true iff reached final token)
```

## Testing
Benchmark test compares `smparser.MatchString` method with `regexp.MatchString`. The input is `https://www.examples.com/tutorials/hello-world/result.json` and given pattern is `https://www*.examples.com/*/*/result.jso*` where '\*' represents wildcard ('.\*' for regexp). The expected evaluation is true. regexp takes an average computation time of 9896 ns/op while smparser computes in 125 ns/op. regexp's memory use for the given example is 8373 B/op and 50 allocs/op. For any given input, smparser's memory use is 0B and 0 alloc.


