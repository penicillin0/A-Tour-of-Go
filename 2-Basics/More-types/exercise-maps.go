package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// なんかノリで書いたらいけた

func WordCount(s string) map[string]int {
	ret_map := make(map[string]int)
	for _, value := range strings.Fields(s) {
		if ret_map[value] != 0 {
			ret_map[value] += 1
		} else {
			ret_map[value] = 1
		}
	}
	return ret_map
}

func main() {
	wc.Test(WordCount)
}
