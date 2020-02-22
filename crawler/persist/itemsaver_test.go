package persist

import (
	"context"
	"eg/egolang/crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name: "AAAA",
		Age:  50,
	}

	id, err := save(expected)

	if err != nil {
		panic(err)
	}

	// TODO: using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhennai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expect %v", actual, expected)
	}
}
