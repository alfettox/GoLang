/*
Author: Giovanni De Franceschi
*/

package main

import (
	"testing"
	"fmt"
)

func Square(x int) int {
	return x * x
}

func TestSquare(t *testing.T) {
	answer := Square(2)
	if answer != 4 {
		t.Errorf("Square(2) did not return 4. Got %d", answer)
	}
}

func TestSquareTableDriven(t *testing.T) {
	var tests = []struct {
		a      int
		expect int
	}{
		{2, 4},
		{0, 0},
		{-3, 9},
		{999999, 999998000001},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.a)
		t.Run(testName, func(t *testing.T) {
			ans := Square(tt.a)
			if ans != tt.expect {
				t.Errorf("got %d, want %d", ans, tt.expect)
			}
		})
	}
}
