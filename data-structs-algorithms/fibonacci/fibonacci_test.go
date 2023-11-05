package main

import "testing"

func Test_fibOneVersion(t *testing.T) {
	actual := fibOneVersion(6)
	if actual != 8 {
		t.Error("expected value 8 in position 6")
	}
}
