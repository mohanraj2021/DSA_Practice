package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		key := getFrequency(str)
		anagramMap[key] = append(anagramMap[key], str)
	}

	var result [][]string
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func getFrequency(str string) string {
	var freq = [26]int{}
	for _, char := range str {
		freq[char-'a']++
	}
	return fmt.Sprint(freq)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
