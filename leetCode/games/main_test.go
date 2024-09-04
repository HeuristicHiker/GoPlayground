package main

import (
	"testing"
)

func TestCalcPoints(t *testing.T) {
	testCase := []string{"5", "2", "C", "D", "+"}
	result := calPoints(testCase)
	soln := 30

	if result != soln {
		t.Errorf("Expected %v but got %v", soln, result)
	}

}
