package main

import (
	"fmt"
	"testing"
)

func intMin(n, m int) int {
	if n < m {
		return n
	}
	return m
}

func TestIntMinBasic(t *testing.T) {
	ans := intMin(1, -2)
	if ans != -2 {
		t.Errorf("intMin(1, -2) = %d; expected %d\n", ans, -2)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		n, m     int
		expected int
	}{
		{0, 1, 0},
		{-1, -2, -2},
		{-2, -1, -2},
		{0, -2, -2},
		{-2, -2, -2},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.n, tt.m)
		t.Run(testname, func(t *testing.T) {
			ans := intMin(tt.n, tt.m)
			if ans != tt.expected {
				t.Errorf("Expected: %d, Got: %d\n", tt.expected, ans)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intMin(1, 2)
	}
}