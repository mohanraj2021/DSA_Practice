package main

import (
	"fmt"
	"math"
)

func main() {
	findTargetArray := []int{2, 7, 11, 15}
	findUniqeuArray := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(findTarget(findTargetArray, 9))
	fmt.Println(findUnique(findUniqeuArray))
	fmt.Println(returnSorterSqaure([]int{-3, -1, 0, 1, 2}))
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

func returnSorterSqaure(arr []int) []int {
	j := 1
	resArr := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		x := arr[i] * arr[i]
		y := arr[j-1] * arr[j-1]
		if x != y {
			resArr[i] = int(math.Max(float64(x), float64(y)))
			j++
		}
	}
	return resArr
}
