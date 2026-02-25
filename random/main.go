package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	ch := make(chan string)
	ch1 := make(chan string)
	for _, str := range strs {
		go getFrequency(str, ch, ch1)
	}
	for i := 0; i < len(strs); i++ {
		k := <-ch
		anagramMap[k] = append(anagramMap[k], <-ch1)
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func getFrequency(str string, ch, ch1 chan string) {
	var freq = [26]int{}
	for _, char := range str {
		freq[char-'a']++
	}
	ch <- fmt.Sprint(freq)
	ch1 <- str
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 4, 5, 4, 4, 6, 5, 5, 1, 3}, 3))
}

func topKFrequent(nums []int, k int) []int {

	result := make([]int, 0)
	var bucket = make([][]int, len(nums))
	var mp = make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for k, v := range mp {
		bucket[v] = append(bucket[v], k)
	}

	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
		}
		if len(result) >= k {
			break
		}
	}
	return result
}
