package main

import "slices"

func SortAndMerge(left, right []int) []int {
	slices.Sort(left)
	slices.Sort(right)

	i, j := 0, 0
	len_l, len_r := len(left), len(right)

	res := make([]int, 0, len_l + len_r)
	for i < len_l && j < len_r {
		if left[i] < right[j] {
			res = append(res, left[i])
			i ++
			continue
		}
		res = append(res, right[j])
		j ++
	}
    res = append(res, left[i:]...)
    res = append(res, right[j:]...)
	return res
}