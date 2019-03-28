package dna

import ("errors")

type Histogram map[rune] int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, char := range d {
		_, ok := h[char]
		if ok {
			h[char] += 1
		} else {
			return h, errors.New("Error")
		}
	}
	return h, nil
}
