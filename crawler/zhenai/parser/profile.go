package parser

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = `<div[^>]*>([0-9]*)岁</div>`
var nicknameRe = `<h1 class="nickName"[^>]*>([^<]*)</h1>`
var idRe = `http://album.zhenai.com/u/([0-9]+)`

var ageCompile = regexp.MustCompile(ageRe)
var nicknameCompile = regexp.MustCompile(nicknameRe)
var idCompile = regexp.MustCompile(idRe)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	match := ageCompile.FindSubmatch(contents)
	if match != nil {
		age, err := strconv.Atoi(extractString(contents, ageCompile))
		if err == nil {
			profile.Age = age
		}
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idCompile),
				Payload: profile,
			},
		},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func ProfileParser(name string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
