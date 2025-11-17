package main

import (
	"testing"
)

// TestAlwaysFails is designed to always fail
func TestAlwaysFails(t *testing.T) {
	// This test will always fail with a clear message
	t.Error("🚨 THIS IS A DELIBERATE TEST FAILURE - This confirms the CI pipeline fails on failing tests")
}

func TestFailingEquality(t *testing.T) {
	got := 2 + 2
	want := 5 // Wrong expectation to force failure

	if got != want {
		t.Fatalf("Deliberate failure: 2+2 should not equal 5. Got %d, want %d", got, want)
	}
}

// TestFailingWithTableDriven tests multiple failing cases
func TestFailingWithTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"Failing case 1", 1, 1, 3}, // 1+1 should be 2, not 3
		{"Failing case 2", 2, 2, 5}, // 2+2 should be 4, not 5
		{"Failing case 3", 3, 3, 7}, // 3+3 should be 6, not 7
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a + tt.b
			if got != tt.want {
				t.Errorf("Addition failed: %d + %d = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
