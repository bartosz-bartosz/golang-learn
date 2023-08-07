package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divideNums(divideNums(tt.divident, tt.divisor))
		if tt.isErr {
			if err == nil {
				t.Error("Expected error but didn't get one")
			}
		} else {
			if err != nil {
				t.Error("Got error, but didn't expect one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDivideNums(t *testing.T) {
	_, err := divideNums(10.0, 1.0)
	if err != nil {
		t.Error("Got an error!")
	}
}

func TestBadDivideNums(t *testing.T) {
	_, err := divideNums(10.0, 0)
	if err != nil {
		t.Error("No error, when it should show up.")
	}
}
