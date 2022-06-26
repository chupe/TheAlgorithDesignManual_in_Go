package robotTourOptimization

import "testing"

func TestWalkUnoptimized(t *testing.T) {
	for _, tc := range testCasesUnoptimized {
		got := NearestNeighbor(tc.points)
		if got != tc.want {
			t.Fatalf("NearestNeighbor(%d) = %f, want %f.",
				tc.points, got, tc.want)
		}
	}
}

func BenchmarkWalkUnoptimized(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesUnoptimized {
			// ignoring errors and results because we're just timing function execution
			_ = NearestNeighbor(tc.points)
		}
	}
}
