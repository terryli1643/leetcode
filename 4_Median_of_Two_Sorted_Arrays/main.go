package main

import "fmt"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := (len(nums1) + len(nums2))
	var sorted []int
	i, j := 0, 0
	for n := 0; n < int(size/2)+1; n++ {
		if i >= len(nums1) {
			sorted = append(sorted, nums2[j])
			j++
		} else if j >= len(nums2) {
			sorted = append(sorted, nums1[i])
			i++
		} else if nums1[i] < nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	if size&1 == 1 {
		return float64(sorted[len(sorted)-1])
	} else {
		return float64(sorted[len(sorted)-2]+sorted[len(sorted)-1]) / 2
	}
}

func main() {
	result := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(result)
}
