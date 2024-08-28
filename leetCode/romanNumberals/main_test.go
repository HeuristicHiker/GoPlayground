package main

import (
	"testing"
)

func TestEight(t *testing.T) {
	result := intToRomanNumeral(8)
	if result != "VIII" {
		t.Errorf("Expect %v but got %v", "VIII", result)
	}
}

func TestNine(t *testing.T) {
	result := intToRomanNumeral(9)
	correct := "IX"
	if result != correct {
		t.Errorf("Expect %v but got %v", correct, result)
	}
}

func TestForty(t *testing.T) {
	result := intToRomanNumeral(40)
	if result != "XL" {
		t.Errorf("Expect %v but got %v", "XL", result)
	}
}
