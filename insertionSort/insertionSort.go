package insertionSort

import "reflect"

func Sort(s []string) []string {
	for i, j := 1, 0; i < len(s); i++ {
		j = i
		for (j > 0) && (s[j] < s[j-1]) {
			swap := reflect.Swapper(s)
			swap(j, j-1)
			j = j - 1
		}
	}

	return s
}
