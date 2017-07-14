package main

import "testing"


func TestAnswer(t *testing.T) {
	a := answer()
	if a != 42 {
		t.Errorf("answer %d != 42", a)
	}
}

func TestHalfAnswer(t *testing.T) {
	a := halfAnswer()
	if a != 21 {
		t.Errorf("halfAnswer %d != 21", a)
	}
}
