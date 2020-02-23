package main

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/model"
	"eg/egolang/crawler_distributed/config"
	"eg/egolang/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {

	go serveRpc(":1234", "test1")
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(":1234")
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1089036739",
		Type: "zhennai",
		Id:   "108906739",
		Payload: model.Profile{
			Name: "AAAA",
			Age:  50,
		},
	}

	result := ""
	client.Call(config.ItemSaverRpc, item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
