package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {

	var arr [3]int
	t.Log(arr[1], arr[2])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 6, 7, 9}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}
