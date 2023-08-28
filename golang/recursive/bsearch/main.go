package main

import "fmt"

var log = fmt.Println

func bsearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var m int
	for l < r {
		m = (l + r) / 2
		if target == nums[m] {
			return m
		}
		if target > nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func recurBsearch(nums []int, t, l, r int) int {
	m := (l + r) / 2
	if nums[m] == t {
		return m
	}
	if l == r && nums[m] != nums[l] {
		return -1
	}
	if t > nums[l] {
		l = m + 1
	} else {
		r = m
	}
	return recurBsearch(nums, t, l, r)
}

func recurBsearch1(nums []int, t, l, r int) int {
	m := (l + r) / 2
	if l > r {
		return -1
	}
	if nums[m] == t {
		return m
	}
	if t > nums[m] {
		return recurBsearch1(nums, t, m+1, r)
	} else {
		return recurBsearch1(nums, t, l, m-1)
	}
}

func RecurBsearch(nums []int, target int) int {
	return recurBsearch1(nums, target, 0, len(nums)-1)
}

func main() {
	nums := []int{1, 3, 7, 9, 10}
	// nums := []int{5}
	// index := bsearch(nums, 5)
	index := RecurBsearch(nums, 1)
	if index == -1 {
		log("not found")
		return
	}
	log("found : ", index)
}
