package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) < 1 {
		return nil
	}

	results := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := results[nums[i]]; ok {
			return []int{results[nums[i]], i}
		}
		results[target-nums[i]] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
