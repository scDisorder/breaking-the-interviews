package codewars

import "strings"

func DNAStrand(dna string) string {
	replacer := strings.NewReplacer(
		"A", "T",
		"T", "A",
		"G", "C",
		"C", "G",
	)
	return replacer.Replace(dna)
}

// TowerBuilder build a tower with specified amount of floors using asterisks and spaces
// example: TowerBuilder(3) will return an array []string{"  *  ", " *** ", "******"}
// to be printed as
// "  *  "
// " *** "
// "*****"
// Source: https://www.codewars.com/kata/576757b1df89ecf5bd00073b
func TowerBuilder(nFloors int) []string {
	if nFloors == 1 {
		return []string{"*"}
	}

	res := make([]string, nFloors)
	for i := 0; i < nFloors; i++ {
		totalAsterisks := (nFloors*2 - 1) - i*2
		spaces := strings.Repeat("+", i)
		b := spaces + strings.Repeat("*", totalAsterisks) + spaces
		res[nFloors-i-1] = b
	}
	return res
}
