package main

import (
	"fmt"
	"strings"
)

func wordfrequency(str string) map[string]int {
	sentence := strings.Fields(str)
	words := make(map[string]int)

	for _, word := range sentence {
		words[word] = words[word] + 1
	}

	return words

}
func main() {
	sentence := "The day at the beach was fun fun fun"
	for key, value := range wordfrequency(sentence) {
		fmt.Println(key, "--", value)
		max := 0
		if value > max {
			max = value
			fmt.Println(key, "word has highest frequency of", max)
		}
	}
}
