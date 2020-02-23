package main

import (
	"eg/egolang/crawler_distributed/config"
	"eg/egolang/crawler_distributed/rpcsupport"
	"eg/egolang/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1238808095",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "天空",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
