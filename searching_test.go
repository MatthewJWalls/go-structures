package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if binarySearch(test, 8) != 8 {
		t.Error("Binary search failed.")
	}

	if binarySearch(test, 9) != 9 {
		t.Error("Binary search failed.")
	}

	if binarySearch(test, 1) != 1 {
		t.Error("Binary search failed.")
	}

	if binarySearch(test, 100) != -1 {
		t.Error("Binary search failed.")
	}

}

func TestNonRecursiveBinarySearch(t *testing.T) {

	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if nonRecursiveBinarySearch(test, 8) != 8 {
		t.Error("Binary search failed.")
	}

	if nonRecursiveBinarySearch(test, 9) != 9 {
		t.Error("Binary search failed.")
	}

	if nonRecursiveBinarySearch(test, 1) != 1 {
		t.Error("Binary search failed.")
	}

	if nonRecursiveBinarySearch(test, 100) != -1 {
		t.Error("Binary search failed.")
	}	

}
