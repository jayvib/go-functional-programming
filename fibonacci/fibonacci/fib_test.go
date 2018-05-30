package fibonacci

import "testing"

var fibTest = []struct {
	a        int
	expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{20, 10946},
	{42, 433494437},
}

func TestFibSimple(t *testing.T) {
	for _, ft := range fibTest {
		if f := fib(ft.a); f != ft.expected {
			t.Errorf("fib(%d) returned %d, expected %d\n", ft.a, f, ft.expected)
		}
	}
}

func BenchmarkFibSimple(b *testing.B) {
	fn := fib
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func TestFibMemoized(t *testing.T) {
	for _, ft := range fibTest {
		if f := FibMemoized(ft.a); f != ft.expected {
			t.Errorf("FibMemoized(%d) returned %d, expected %d\n", ft.a, f, ft.expected)
		}
	}
}

func BenchmarkFibMemoized(b *testing.B) {
	fn := FibMemoized
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func TestFibChanneled(t *testing.T) {
	for _, ft := range fibTest {
		if f := FibChanneled(ft.a); f != ft.expected {
			t.Errorf("FibChanneled(%d) returned %d, expected %d\n", ft.a, f, ft.expected)
		}
	}
}

func BenchmarkFibChanneled(b *testing.B) {
	fn := FibChanneled
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}
