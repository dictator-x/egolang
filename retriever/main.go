package main

import (
	"eg/egolang/retriever/mock"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"fake content"}
	fmt.Println(download(r))
}
