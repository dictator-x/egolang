package main

import (
	"eg/egolang/retriever/mock"
	"eg/egolang/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}

}

func main() {
	var r Retriever

	r = mock.Retriever{"this is fack"}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.TimeOut)
	}
	// fmt.Println(download(r))
}
