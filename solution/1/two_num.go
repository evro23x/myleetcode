package main

import "fmt"

func twoSum(nums []int, target int) []int {
		seenNums := make(map[int]int)
		for index, thisNum := range nums {
			if seenIndex, ok := seenNums[target - thisNum]; ok {
				return []int{seenIndex, index}
			}
			seenNums[thisNum] = index
		}
	return []int{}
}

func main() {
	nums := []int{3,2,3}
	target := 9
	twoSum(nums, target)
	fmt.Println("done")
}

//func twoSum(nums []int, target int) []int {
//	seenNums := make(map[int]int)
//	for index, thisNum := range nums {
//		if seenIndex, ok := seenNums[target - thisNum]; ok {
//			return []int{seenIndex, index}
//		}
//		seenNums[thisNum] = index
//	}
//	return []int{0, 0} // Should not happen
//}


//func twoSum(nums []int, target int) []int {
//	lookup := make(map[int]int)
//	for i, v := range nums {
//		j, ok := lookup[-v]
//		lookup[v - target] = i
//		if ok {
//			return []int{j, i}
//		}
//	}
//	return []int{}
//}

//func twoSum(nums []int, target int) []int {
//	numsMap := make(map[int]int, len(nums))
//	for i, v := range nums {
//		numsMap[v] = i
//	}
//	for i, v := range nums {
//		if j, ok := numsMap[target - v]; ok && j != i {
//			return []int{i, j}
//		}
//	}
//	return []int{}
//}