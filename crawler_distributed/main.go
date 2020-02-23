package main

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/scheduler"
	"eg/egolang/crawler/zhenai/parser"
	"eg/egolang/crawler_distributed/config"
	itemsaver "eg/egolang/crawler_distributed/persist/client"
	worker "eg/egolang/crawler_distributed/worker/client"
	"fmt"
)

func main() {
	// resp, err := http.Get("http://www.zhenai.com/zhenghun")

	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })

	// fmt.Printf("%s\n", all)
	// printCityList(all)

}
