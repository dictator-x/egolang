package main

import (
	"fmt"
)

func lengthOfNoRepeatingSubStr(s string) int {
	// lastOccurred := make(map[rune]int)
	lastOccurred := make([]int, 0xffff)

	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength

}

func main() {
	fmt.Println(lengthOfNoRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNoRepeatingSubStr(""))
	fmt.Println(lengthOfNoRepeatingSubStr("b"))
}
