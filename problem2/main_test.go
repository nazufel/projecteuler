package main

import "testing"

type testFibPairs struct {
	term     int
	expected int
}

var fibPairs = []testFibPairs{
	{10, 55},
	{20, 6765},
	{31, 1346269},
}

func TestFib(t *testing.T) {
	for _, pair := range fibPairs {
		v := FibonacciLoop(pair.term)
		if v != pair.expected {
			t.Error(
				"For", pair.term,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}
