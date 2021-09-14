package utils

import "testing"

func TestSum(t *testing.T) {
	a := 1
	b := 2

	c := Sum(a, b)

	if (c != 3) {
		t.Errorf("expected %v but go %v", 3, c)
	}
}