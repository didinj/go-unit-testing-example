package math

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add two positives", 2, 3, 5},
		{"Add positive and negative", 5, -2, 3},
		{"Add zeros", 0, 0, 0},
		{"Add two negatives", -1, -1, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		want        float64
		expectError bool
	}{
		{"Valid division", 10, 2, 5, false},
		{"Division by zero", 5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Fatal("expected error, got none")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if got != tt.want {
					t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
				}
			}
		})
	}
}
