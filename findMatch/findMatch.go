package findMatch

import "unicode/utf8"

func FindMatch(s1, s2 string) int {
	for i := range s1 {
		for j, c2 := range s2 {
			if i+j >= utf8.RuneCountInString(s1) {
				break
			}
			if c2 != []rune(s1)[i+j] {
				break
			}
			if j == utf8.RuneCountInString(s2)-1 {
				return i
			}
		}
	}
	return -1
}
