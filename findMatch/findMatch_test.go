package findMatch

import "testing"

func TestSorting(t *testing.T) {
	for _, tc := range testCases {
		got := FindMatch(tc.s1, tc.s2)
		if got != tc.want {
			t.Fatalf("FindMatch(%s, %s) = %d, want %d.",
				tc.s1, tc.s2, got, tc.want)
		}
	}
}

func BenchmarkSorting(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			_ = FindMatch(tc.s1, tc.s2)
		}
	}
}
