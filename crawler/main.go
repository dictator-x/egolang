package main

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/persist"
	"eg/egolang/crawler/scheduler"
	"eg/egolang/crawler/zhenai/parser"
)

func main() {
	// resp, err := http.Get("http://www.zhenai.com/zhenghun")

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})

	// fmt.Printf("%s\n", all)
	// printCityList(all)

}
