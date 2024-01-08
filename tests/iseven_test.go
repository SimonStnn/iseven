package iseven_test

import (
	. "github.com/SimonStnn/iseven"
	"testing"
)

func TestIsEven(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{-5, false},
		{-4, true},
		{-3, false},
		{-2, true},
		{-1, false},
		{0, true},
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
	}

	for _, tt := range tests {
		if got := IsEven(tt.input); got != tt.want {
			t.Errorf("IsOdd(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
