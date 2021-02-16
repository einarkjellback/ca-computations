package cacomp

import "testing"

func TestUpdate(t *testing.T) {
	var r Rules = 1 + 4 + 32
	c := Config{false, true, false, false, true, true}
	expect := Config{true, true, false, false, false, false}
	actual, err := Update(c, 1, r)
	if err != nil {
		t.Fatal(err)
	}
	isEqual := len(expect) == len(actual)
	for i := 0; isEqual && i < len(expect); i++ {
		isEqual = expect[i] == actual[i]
	}
	if !isEqual {
		t.Errorf("expected %v, but was %v", expect, actual)
	}
}
