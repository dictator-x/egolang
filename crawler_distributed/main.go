package main

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/scheduler"
	"eg/egolang/crawler/zhenai/parser"
	"eg/egolang/crawler_distributed/config"
	"eg/egolang/crawler_distributed/persist/client"
	"fmt"
)

func main() {
	// resp, err := http.Get("http://www.zhenai.com/zhenghun")

	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
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
