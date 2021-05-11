package odd_counts_character

import "strings"

func OddCountsCharacter(n int) string {
	if n%2 == 0 {
		return strings.Repeat("a", n-1) + "b"
	}

	return strings.Repeat("a", n)
}
