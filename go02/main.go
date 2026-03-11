package main

import (
	"sort"
)

func main() {
}

func GroupAnagrams(words []string) map[string][]string {
	anagram := make(map[string][]string)
	for _, word := range words {
		soredWord := sortStrings(word)
		anagram[soredWord] = append(anagram[soredWord], word)
	}
	return anagram
}

func sortStrings(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
