package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := arr[2:6]
	fmt.Println(s)
}
