package anagrams

func Anagrams(s1, s2 string) bool {
	m := make(map[rune]int)

	for _, c := range s1 {
		if _, exists := m[c]; !exists {
			m[c] = 0
		}
		m[c] += 1
	}

	for _, c := range s2 {
		if _, exists := m[c]; exists {
			m[c] -= 1
		} else {
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
