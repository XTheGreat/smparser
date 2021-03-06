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
```
BenchmarkSmparser-8   	 9024852	       135 ns/op	       0 B/op	       0
BenchmarkRegexp-8   	  122932	      9661 ns/op	    8363 B/op	      50 allocs/op
```


