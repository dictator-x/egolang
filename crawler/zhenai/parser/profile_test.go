package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/1089036739", "nickname")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
}
