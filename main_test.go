package main

import "testing"

func TestSubtract(t *testing.T) {
	res := Subtract(1200, 780)
	if res != 420 {
		t.Error()
	}
}
