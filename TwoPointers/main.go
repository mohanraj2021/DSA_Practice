package main

import "fmt"

func main() {
	findTargetArray := []int{2, 7, 11, 15}
	findUniqeuArray := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(findTarget(findTargetArray, 9))
	fmt.Println(findUnique(findUniqeuArray))
}

func findTarget(arr []int, target int) []int {
	left, right := 0, len(arr)-1
	var res []int
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			res = append(res, left, right)
			fmt.Println(res)
		}
		if sum > target {
			right = right - 1

		} else {
			left++
		}
	}
	return res
}

func findUnique(arr []int) []int {

	j := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j-1] {
			arr[j] = arr[i]
			j++
		}
	}
	return arr[:j]
}
