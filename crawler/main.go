package main

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/persist"
	"eg/egolang/crawler/scheduler"
	"eg/egolang/crawler/zhenai/parser"
)

func main() {
	// resp, err := http.Get("http://www.zhenai.com/zhenghun")

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

	// e.Run(engine.Request{
	// 	Url:    "http://www.zhenai.com/zhenghun/shanghai",
	// 	Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
	// })

	// fmt.Printf("%s\n", all)
	// printCityList(all)

}
