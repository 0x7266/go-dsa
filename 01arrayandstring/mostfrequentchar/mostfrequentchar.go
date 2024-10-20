package mostfrequentchar

import (
	"fmt"
	"math"
)

func MostFrequentChar(s string) rune {
	m := make(map[rune]int)
	var c rune
	max := int(math.Inf(-1))
	fmt.Println(max)

	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = 0
		}

		m[c] += 1
	}

	for k, v := range m {
		if v > max {
			max = v
			c = k
		}
	}

	return c
}
