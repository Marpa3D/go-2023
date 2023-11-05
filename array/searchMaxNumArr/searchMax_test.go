package main

import "testing"

func Test_searchMax(t *testing.T) {
	arr := []int{2, 34, 8, 24}
	got := searchMax(arr)
	if got != 34 {
		t.Error("Ожидается число 34!")
	}
}
