package parser

import (
	"eg/egolang/crawler/fetcher"
	"eg/egolang/util/file"
	// "fmt"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", contents)

	file.TruncWriteFile("cityList.text", string(contents))
}
