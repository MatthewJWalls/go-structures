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

	log.Print("ok now looking at: ", arr, "with length ", len(arr))

	if len(arr) > 1 {
		log.Print("  Gonna tackle it because > 1")
	}

	if len(arr) > 1 {

		pivot  := arr[len(arr)/2]
		left  := 0
		right := len(arr)-1

		log.Printf("  Using pivot %d", pivot)

		for left <= right  {

			for arr[left] < pivot {
				left++
			} 

			for arr[right] > pivot {
				right--
			}

			if left <= right {
				log.Printf("  Swapping %d and %d", arr[left], arr[right])
				tmp := arr[left]
				arr[left] = arr[right]
				arr[right] = tmp
				left++
				right--
			}

		}

		log.Print("  Finished swaps.")
		pindex := len(arr)/2
		log.Print("  sorting ", arr[:pindex], " and ", arr[pindex:] )
		lefter  := quicksort(arr[:pindex])
		righter := quicksort(arr[pindex:])

		log.Print("  appending ", lefter, " and", righter)
		log.Print("giving ", append(lefter, righter...))
		return append(lefter, righter...)

	} else {

		log.Print("  skipped it")
		return arr

	}

}

