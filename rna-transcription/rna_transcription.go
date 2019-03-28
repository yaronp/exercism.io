package strand
import "bytes"

func ToRNA(dna string) string {
	r := map[rune]rune{'G': 'C', 'C' : 'G', 'T' : 'A',	'A' : 'U' }
	var result bytes.Buffer

	for _, s := range dna {
		result.WriteRune(r[s])
	}
	return result.String()
}
