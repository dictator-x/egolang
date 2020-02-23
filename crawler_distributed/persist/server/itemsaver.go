package main

import (
	"eg/egolang/crawler_distributed/config"
	"eg/egolang/crawler_distributed/persist"
	"eg/egolang/crawler_distributed/rpcsupport"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(":1234",
		&persist.ItemSaverService{
			Client: client,
			Index:  "dating_profile",
		})
}
