package main

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
