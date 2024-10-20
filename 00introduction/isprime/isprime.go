package isprime

import (
	"fmt"
	"math"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n == 9 {
			fmt.Println(i)
		}
		if n%i == 0 {
			return false
		}
	}
	return true
}
