package main

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := cache[target-nums[i]]; ok {
			return []int{v, i}
		}
		cache[nums[i]] = i
	}
	return nil
}
