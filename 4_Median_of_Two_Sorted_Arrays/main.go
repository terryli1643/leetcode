package main

import "fmt"

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)

	p1, p2 := 0, 0
	if size%2 == 1 {
		for i := 0; i < size/2-1; i++ {
			if p2 == len(nums2)-1 {
				p1++
			} else if p1 == len(nums1)-1 {
				p2++
			} else if nums1[p1] > nums2[p2] {
				p2++
			} else {
				p1++
			}
		}
		if nums1[p1] > nums2[p2] {
			return float64(nums1[p1])
		} else {
			return float64(nums2[p2])
		}
	} else {
		n1, n2 := 0, 0
		for i := 0; i < size/2; i++ {
			if nums1[p1] > nums2[p2] {
				p2++
			} else {
				p1++
			}
		}
		return float64((n1 + n2) / 2)
	}
}

func main() {
	result := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(result)
}
