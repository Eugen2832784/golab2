// Jevgēnijs Kovaļonoks 251RDB125
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Print("First word: ")
	word1 := sanemtVardu()
	fmt.Print("Second word: ")
	word2 := sanemtVardu()
	fmt.Println(irAnagrammas(word1, word2))
}

func irAnagrammas(vards1, vards2 string) string {
	if utf8.RuneCountInString(vards1) != utf8.RuneCountInString(vards2) {
		return "NO"
	}

	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, char := range vards1 {
		map1[char]++
	}
	for _, char := range vards2 {
		map2[char]++
	}

	for i, count := range map1 {
		if map2[i] != count {
			return "NO"
		}
	}

	return "YES"
}

func sanemtVardu() string {
	var word string
	fmt.Scanln(&word)
	word = strings.ToLower(word)
	return word
}
