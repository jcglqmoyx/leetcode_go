package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for idx := m + n - 1; idx >= 0; idx-- {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[idx] = nums1[i]
				i--
			} else {
				nums1[idx] = nums2[j]
				j--
			}
		} else if i >= 0 {
			nums1[idx] = nums1[i]
			i--
		} else {
			nums1[idx] = nums2[j]
			j--
		}
	}
}
