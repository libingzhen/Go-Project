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
	var lowIndex int
	var highIndex = len(nums) - 1
	for lowIndex < highIndex {
		var midIndex = (highIndex + lowIndex) // 2
		if target > nums[midIndex] {
			lowIndex = midIndex + 1
		} else if target < nums[midIndex] {
			highIndex = midIndex - 1
		} else {
			return nums[midIndex]
		}
	}
	return 0
}

func selectionSort(nums []int) []int {
	var newArr []int
	var capacity = len(nums)
	for i := 0; i < capacity; i++ {
		var eachInddex = 0
		var smallest = nums[0]
		for j := 0; j < len(nums); j++ {
			if smallest > nums[j] {
				smallest = nums[j]
				eachInddex = j
			}
		}
		if len(nums) > 1 {
			if eachInddex == 0 {
				nums = append(nums[1:])
			} else {
				nums = append(nums[:eachInddex], nums[eachInddex+1:]...)
			}
			newArr = append(newArr, smallest)
		} else {
			newArr = append(newArr, smallest)
		}
	}
	return newArr
}
