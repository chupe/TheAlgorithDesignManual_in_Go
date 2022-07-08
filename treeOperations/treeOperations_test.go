package treeOperations

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCompare(t *testing.T) {
	for _, tc := range compareTestCases {
		got := tc.t1.Compare(tc.t2)
		if got != tc.want {
			t.Fatalf("Compare(%s\n, %s\n) = %+v\n, want %+v.",
				tc.t1, tc.t2, got, tc.want)
		}
	}
}

func TestInsert(t *testing.T) {
	for _, tc := range insertTestCases {
		got := Insert(tc.t, tc.v)
		if !got.Compare(tc.want) {
			t.Fatalf("Insert(%+v\n, %d) = %+v\n, want %+v\n.",
				tc.t, tc.v, got, tc.want)
		}
	}
}

func TestSearch(t *testing.T) {
	for _, tc := range searchTestCases {
		got := Search(tc.t, tc.v)
		if !got.Compare(tc.want) {
			t.Fatalf("Search(%s, %d) = %+v\n, want %+v\n.",
				tc.t, tc.v, got, tc.want)
		}
	}
}

func TestMin(t *testing.T) {
	for _, tc := range minTestCases {
		got := Min(tc.t)
		if !got.Compare(tc.want) {
			t.Fatalf("Min(%s) = %+v\n, want %+v\n.",
				tc.t, got, tc.want)
		}
	}
}

func TestTraverseTree(t *testing.T) {
	for _, tc := range traverseTestCases {
		got := TraverseTree(tc.t)
		if !cmp.Equal(got, tc.want) {
			t.Fatalf("TraverseTree(%s) = %+v\n, want %+v\n.",
				tc.t, got, tc.want)
		}
	}
}

func TestDelete(t *testing.T) {
	for _, tc := range deleteTestCases {
		got := Delete(tc.t, tc.v)
		if !got.Compare(tc.want) {
			t.Fatalf("Delete(%+v\n, %d) = %+v\n, want %+v\n.",
				tc.t, tc.v, got, tc.want)
		}
	}
}

func TestFindParent(t *testing.T) {
	for _, tc := range findParentTestCases {
		got := tc.c.Parent(tc.p)
		if !got.Compare(tc.want) {
			t.Fatalf("%s.Parent(%s) = %+v\n, want %+v\n.",
				tc.c, tc.p, got, tc.want)
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range insertTestCases {
			_ = Insert(tc.t, tc.v)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range searchTestCases {
			_ = Search(tc.t, tc.v)
		}
	}
}

func BenchmarkMin(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range minTestCases {
			_ = Min(tc.t)
		}
	}
}

func BenchmarkTraverseTree(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range traverseTestCases {
			_ = TraverseTree(tc.t)
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range deleteTestCases {
			_ = Delete(tc.t, tc.v)
		}
	}
}
