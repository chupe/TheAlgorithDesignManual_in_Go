package insertionSort

import "testing"

func TestSorting(t *testing.T) {
	for _, tc := range testCases {
		got := Sort(tc.s1)
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("Sort(%s) = %s, want %s.",
					tc.s1, got, tc.want)
			}
		}
	}
}

func BenchmarkSorting(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			_ = Sort(tc.s1)
		}
	}
}
