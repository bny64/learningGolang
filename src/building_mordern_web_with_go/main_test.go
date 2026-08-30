package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, test := range tests {
		got, err := divide(test.dividend, test.divisor)

		if test.isErr {
			if err == nil {
				t.Error("ecpected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Expected no error, but got", err.Error())
			}
			if got != test.expected {
				t.Error("Expected", test.expected, "but got", got)
			}
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Expected no error, but got", err)
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
