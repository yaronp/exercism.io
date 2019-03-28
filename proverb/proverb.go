package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	var length = len(rhyme)

	if length == 0 {
		return []string{}
	}

	const (
		tail = "And all for the want of a %s."
		statement = "For want of a %s the %s was lost."
	)

	var slice = make ([]string, length)

	for i := 0; i < length -1; i++ {
		slice[i] =  fmt.Sprintf(statement, rhyme[i], rhyme[i+1])
	}

	slice[length-1] = fmt.Sprintf(tail, rhyme[0])
	return slice
}
