package intutils

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(1, 2)
	if ans != 1 {
		t.Errorf("IntMin(1, 2) wanted 1 got %v", ans)
	}
}

func TestIntMinTable(t *testing.T) {
	var tests = []struct {
		a, b     int
		expected int
	}{
		{0, 1, 0},
		{2, -2, -2},
		{2, 7, 2},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.expected {
				t.Errorf("got %v, wanted %v", ans, tt.expected)
			}
		})
	}
}
