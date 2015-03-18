package main

import "log"

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left  := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left []int, right []int) []int {

	out := make([]int, 0, len(left)+len(right))

	for len(left) != 0 && len(right) != 0 {

		if left[0] < right[0] {
			out = append(out, left[0])
			left = left[1:]
		} else {
			out = append(out, right[0])
			right = right[1:]
		}

	}

	if len(left) > 0 {
		out = append(out, left...)
	}
	if len(right) > 0 {
		out = append(out, right...)
	}

	return out

}

func quicksort(arr []int) []int {

	var pivot int

	log.Print(arr)

	if len(arr) > 1 {

		pivot  = arr[len(arr)/2]
		left  := 0
		right := len(arr)-1

		log.Printf("Using pivot %d", pivot)

		for left <= right {

			log.Printf("Looking for a left")

			for arr[left] < pivot {
				left++
			} 

			log.Printf("Looking for a right")

			for arr[right] > pivot {
				right--
			}

			log.Printf("Looking at swapping %d and %d", arr[left], arr[right])

			if left <= right {
				log.Printf("Swapping %d and %d", arr[left], arr[right])
				tmp := arr[left]
				arr[left] = arr[right]
				arr[right] = tmp
				left++
				right++
			} else {
				panic("NO SWAP?")
			}

		}

		lefter  := quicksort(arr[:left])
		righter := quicksort(arr[left:])

		return append(lefter, righter...)

	} else {

		return arr

	}

}













