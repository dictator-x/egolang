package main

import (
	"eg/egolang/crawler_distributed/config"
	"eg/egolang/crawler_distributed/rpcsupport"
	"eg/egolang/crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{},
	))
}
