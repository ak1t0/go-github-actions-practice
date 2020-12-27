package main

import "testing"

func TestDummy(t *testing.T) {
	if r := Dummy(6); r == 36 {
		t.Log("Success")
	} else {
		t.Error("Fail")
	}
}
