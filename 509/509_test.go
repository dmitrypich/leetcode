package main

import "testing"

func TestFib(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
		{16, 987},
		{17, 1597},
		{18, 2584},
		{19, 4181},
	}
	for _, test := range tests {
		if got := fib(test.input); got != test.want {
			t.Errorf("Fib(%d) = %d", test.input, got)
		}
	}
}

func TestFibRe(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
		{14, 377},
		{15, 610},
		{16, 987},
		{17, 1597},
		{18, 2584},
		{19, 4181},
	}
	for _, test := range tests {
		if got := fib(test.input); got != test.want {
			t.Errorf("FibRe(%d) = %d", test.input, got)
		}
	}
}

func BenchmarkFib5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(5)
	}
}

func BenchmarkFib25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(25)
	}
}

func BenchmarkFib50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(50)
	}
}

func BenchmarkFib500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(500)
	}
}

func BenchmarkFibRe5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRe(5)
	}
}

func BenchmarkFibRe10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRe(10)
	}
}

func BenchmarkFibRe25(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRe(25)
	}
}
