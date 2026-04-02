package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("call once", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"
		assertEqual(t, got, want)
	})

	t.Run("call five times", func(t *testing.T) {
		got := Repeat("z", 5)
		want := "zzzzz"
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		Repeat("a", 50)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 3))
	// Output: aaa
}
