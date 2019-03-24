package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 4!!!")
	}
}

func TestCalculateStruct(t *testing.T) {
	var vals = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{6, 8},
		{-2, 0},
	}

	for _, val := range vals {
		// if Calculate(val.input) != val.expected {
		if ouput := Calculate(val.input); ouput != val.expected {
			t.Error(val.input, val.expected, "Uyumsuz!!!")
		}
	}
}
