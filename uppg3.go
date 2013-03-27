package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var res = make(map[string]int)
	var sa = strings.Fields(s)

	for i := range sa {
		elem := res[sa[i]]

		if elem > 0 {
			res[sa[i]]++
		} else {
			res[sa[i]] = 1
		}
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
