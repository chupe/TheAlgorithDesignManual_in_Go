package selectingTheRightJobs

import "testing"

func TestEarliestFirst(t *testing.T) {
	for _, tc := range testCasesEarliest {
		got := EarliestFirst(tc.s1)
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("EarliestFirst(%s) = %s, want %s.",
					tc.s1, got, tc.want)
			}
		}
	}
}

func TestEarliestFirstV2(t *testing.T) {
	for _, tc := range testCasesEarliest {
		got := EarliestFirstV2(tc.s1)
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("EarliestFirstV2(%s) = %s, want %s.",
					tc.s1, got, tc.want)
			}
		}
	}
}

func TestShortestFirst(t *testing.T) {
	for _, tc := range testCasesShortest {
		got := ShortestFirst(tc.s1)
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("ShortestFirst(%s) = %s, want %s.",
					tc.s1, got, tc.want)
			}
		}
	}
}

func TestOptimalScheduling(t *testing.T) {
	for _, tc := range testCasesOptimal {
		got := OptimalScheduling(tc.s1)
		for i := range got {
			if got[i] != tc.want[i] {
				t.Fatalf("OptimalScheduling(%s) = %s, want %s.",
					tc.s1, got, tc.want)
			}
		}
	}
}

func BenchmarkEarliestFirst(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesEarliest {
			_ = EarliestFirst(tc.s1)
		}
	}
}

func BenchmarkEarliestFirstV2(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesEarliest {
			_ = EarliestFirstV2(append([]Job(nil), tc.s1...))
		}
	}
}

func BenchmarkShortestFirst(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesShortest {
			_ = ShortestFirst(tc.s1)
		}
	}
}

func BenchmarkOptimalScheduling(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesOptimal {
			_ = OptimalScheduling(tc.s1)
		}
	}
}
