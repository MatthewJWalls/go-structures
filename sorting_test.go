package main

import (
	"testing"
)

func TestMerge(t *testing.T){

	a := []int{1, 3, 5}
	b := []int{2, 4, 6}

	c := merge(a, b)

	if c[0] != 1 {
		t.Error("Merge doesn't work.")
	}

	if c[1] != 2 {
		t.Error("Merge doesn't work.")
	}

	if c[5] != 6 {
		t.Error("Merge doesn't work.")
	}

}

func TestMergeSort(t *testing.T){

	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	output := mergeSort(input)

	if len(output) != 10 {
		t.Error("Not enough values after merge sort")
	}

	if output[0] != 0 {
		t.Error("Incorrect value after merge sort")
	}

	if output[1] != 1 {
		t.Error("Incorrect value after merge sort")
	}

	if output[9] != 9 {
		t.Error("Incorrect value after merge sort")
	}

	// more test sets

	if mergeSort([]int{4, 5, 6, 1, 2, 3, 0})[3] != 3 {
		t.Error("Incorrect value after merge sort")
	}

	if mergeSort([]int{1, 1, 1, 3, 4, 5, 6, 7, 7, 8})[3] != 3 {
		t.Error("Incorrect value after merge sort")
	}

}

func TestQuickSort(t *testing.T) {

	output := quicksort([]int{5, 0, 3, 2, 1, 6, 6})

	if ! intArraysEqual(output, []int{0, 1, 2, 3, 5, 6, 6}) {
		t.Error("Incorrect value after quick sort.")
	}

	if ! intArraysEqual(quicksort([]int{0, -1, -2, -3}), []int{-3, -2, -1, 0}) {
		t.Error("Incorrect value after quick sort.")
	}

	if ! intArraysEqual(quicksort([]int{1, 1, 1}), []int{1, 1, 1}) {
		t.Error("Incorrect value after quick sort.")
	}

	if ! intArraysEqual(quicksort([]int{6, 1, 1}), []int{1, 1, 6}) {
		t.Error("Incorrect value after quick sort.")
	}

}
