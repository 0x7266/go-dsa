package maxvalue

import "math"

func MaxValue(numbers []float32) float32 {
	max := float32(math.Inf(-1))
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}
