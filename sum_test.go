package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Errorf("The result must be 5")
	}
}

func TestSub(t *testing.T) {

	result := sub(2, 3)

	if result != 1 {
		t.Errorf("The result must be 1")
	}
}

// func TestMult(t *testing.T) {

// 	result := Mult(2, 3)

// 	if result != 12 {
// 		t.Errorf("The result must be 12")
// 	}
// }

func TestTimes(t *testing.T) {

	result := times(2, 3)

	if result != 6 {
		t.Errorf("The result must be 6")
	}
}

func TestSumMult(t *testing.T) {

	result := sumMult(2, 3)

	if result != 10 {
		t.Errorf("The result must be 10")
	}
}
