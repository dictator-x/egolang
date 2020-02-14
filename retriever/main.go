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

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

const url = "http://www.imooc.com"

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "xxx",
		},
	)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "fackkkkk",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}

}

func main() {
	var r Retriever

	rmock := &mock.Retriever{"this is fack"}
	fmt.Printf("%T %v\n", r, r)

	fmt.Println("session")
	fmt.Println(session(rmock))

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
