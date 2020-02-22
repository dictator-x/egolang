package persist

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d: %v", itemCount, item)
			itemCount++
			save(item)
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhennai").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
