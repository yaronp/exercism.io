package hamming

import (
	"errors"
)

// Distance check the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("mismatch (one string too long from other)")
	}

	var diff = 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
