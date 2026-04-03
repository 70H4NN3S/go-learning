package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{123, 234, 345, 456, 654}
	b.ReportAllocs()
	for b.Loop() {
		Sum(numbers)
	}
}
