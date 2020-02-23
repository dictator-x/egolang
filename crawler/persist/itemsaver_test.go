package persist

import (
	"context"
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1089036739",
		Type: "zhennai",
		Id:   "108906739",
		Payload: model.Profile{
			Name: "AAAA",
			Age:  50,
		},
	}

	const index = "dating_test"

	// TODO: using docker go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err)
	}

	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expect %v", actual, expected)
	}
}
