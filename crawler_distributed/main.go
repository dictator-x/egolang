package main

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/scheduler"
	"eg/egolang/crawler/zhenai/parser"
	itemsaver "eg/egolang/crawler_distributed/persist/client"
	"eg/egolang/crawler_distributed/rpcsupport"
	worker "eg/egolang/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	// resp, err := http.Get("http://www.zhenai.com/zhenghun")

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)
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
	// 	Url:    "http://www.zhenai.com/zhenghun/shanghai",
	// 	Parser: engine.NewFuncParser(parser.ParseCity, "ParseCity"),
	// })

	// fmt.Printf("%s\n", all)
	// printCityList(all)

}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out

}
