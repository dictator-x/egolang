package main

import (
	"eg/egolang/crawler_distributed/persist"
	"eg/egolang/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
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
