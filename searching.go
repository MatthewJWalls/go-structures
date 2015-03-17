package main

import "log"

func binarySearch(arr []int, val int) int {

	midpoint := len(arr)/2

	if arr[midpoint] == val {
		return arr[midpoint]
	} else if len(arr) == 1 {
		return -1
	} else if val > arr[midpoint] {
		// recurse right
		return binarySearch(arr[midpoint:], val)
	} else if val < arr[midpoint] {
		// recurse left
		return binarySearch(arr[:midpoint], val)
	} else {
		panic("How did we get here?")
	}

}

func nonRecursiveBinarySearch(arr []int, val int) int {

	midpoint := len(arr)/2
	begin := 0
	end := len(arr)

	for {

		midpoint = begin+((end-begin)/2)

		if arr[midpoint] == val {
			// found it
			return arr[midpoint]
		} else if end-begin == 1 {
			// not there
			return -1
		} else if val > arr[midpoint] {
			// go right
			begin = midpoint
		} else if val < arr[midpoint] {
			// go left
			end = midpoint
		}


	}

}
