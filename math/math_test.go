package math

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	if sum != 6 {
		t.Log("test fail")
		t.Fail()
	}
}
