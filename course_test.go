package main

import "testing"


func TestAnswer(t *testing.T) {
	a := answer()
	if a != 42 {
		t.Errorf("answer %d != 42", a)
	}
}
