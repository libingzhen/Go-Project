package main

func twoSum(nums []int, target int) []int {
	var m, n int
	for m = 0; m < len(nums); m++ {
		for n = 0; n < len(nums); n++ {
			if nums[m]+nums[n] == target && m != n {
				return []int{m, n}
			}
		}
	}
	return nil
}

//func twoSumHash(nums []int, target int) []int{

//}

func binarySearch(nums []int, target int) int {
	var low_index int = 0
	var high_index int = len(nums) - 1
	for low_index < high_index {
		var mid_index int = (high_index + low_index) // 2
		if target > nums[mid_index] {
			low_index = mid_index + 1
		} else if target < nums[mid_index] {
			high_index = mid_index - 1
		} else {
			return nums[mid_index]
		}
	}
	return -1
}

func selectionSort(nums []int, target int) []int {

}
