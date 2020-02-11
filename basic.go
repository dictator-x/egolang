package main

import (
	"fmt"
	"math"
	"runtime"
)

// can not use := outside of function
var aa = 3
var ss = "kkk"

var (
	aa1 = 3
	ss1 = "kkk1"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, s = 3, 4, "def"
	fmt.Println(a, b, s)
}

func variableShorter() {
	a, b, s := 3, 5, "dea"
	fmt.Println(a, b, s)
}

// const can be use by any relative type
func consts() {
	const filename string = "abc.txt"
	const b = 4
	var a = 3
	var c int
	c = int(math.Sqrt(float64(a*a) + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
		javascript
	)
	fmt.Println(cpp, java, python, golang, javascript)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println(runtime.GOARCH)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, aa1, ss, ss1)
	consts()
	enums()
}
