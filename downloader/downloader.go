package main

import (
	// "eg/egolang/infra"
	"eg/egolang/testing"
	"fmt"
)

func getRetriever() retriever {
	return testing.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
