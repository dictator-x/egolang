package main

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/zhenai/parser"
)

func main() {
	// resp, err := http.Get("http://www.zhenai.com/zhenghun")

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	// fmt.Printf("%s\n", all)
	// printCityList(all)

}
