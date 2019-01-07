package sort

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
