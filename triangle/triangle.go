package triangle

import (
	"math"
)

// Kind is triangle type
type Kind int // Triangle type

const (
	// NaT is not a triangle
	NaT = 1
	// Equ equilateral
	Equ = 2
	// Iso isosceles
	Iso = 3
	// Sca scalene
	Sca = 4
)

func checkInvalidNumber(a float64) bool {
	return math.IsInf(a, 1) || math.IsInf(a, -1) || math.IsNaN(a) || a == 0
}

// KindFromSides determinate triangle type.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if checkInvalidNumber(a) || checkInvalidNumber(b) || checkInvalidNumber(c) {
		k = NaT
	} else if a+b < c || a+c < b || b+c < a {
		k = NaT
	} else if a == b && b == c && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else if a != b && a != c && b != c {
		k = Sca
	} else {
		k = NaT
	}
	return k
}
