package main

import "fmt"

func main() {
	words := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}

	fmt.Println(ParseAnagram(words))
}

func ParseAnagram(words []string) [][]string {
	var anagrams [][]string
	var index []int

	for i := 0; i < len(words); i++ {
		anagram := []string{words[i]}

		if CheckIndex(index, i) {
			continue
		}
		
		for j := i + 1; j < len(words); j++ {
			if IsAnagram(words[i], words[j]) {
				index = append(index, j)
				anagram = append(anagram, words[j])
			}
		}
		anagrams = append(anagrams, anagram)
	}
	return anagrams
}

func IsAnagram(word1, word2 string) bool {
	res1 := 0
	for _, char := range word1 {
		res1 += int(char)
	}

	res2 := 0
	for _, char := range word2 {
		res2 += int(char)
	}

	return res1 == res2
}

func CheckIndex(v []int, value int) bool {
	valid := false
	for i := 0; i < len(v); i++ {
		if v[i] == value {
			valid = true
			break
		}
	}
	return valid
}