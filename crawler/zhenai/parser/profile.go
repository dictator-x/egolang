package parser

import (
	"eg/egolang/crawler/engine"
	"eg/egolang/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = `<div[^>]*>([0-9]*)岁</div>`
var nicknameRe = `<h1 class="nickName"[^>]*>([^<]*)</h1>`

var ageCompile = regexp.MustCompile(ageRe)
var nicknameCompile = regexp.MustCompile(nicknameRe)

func ParseProfile(contents []byte) engine.ParseResult {
	match := ageCompile.FindSubmatch(contents)

	profile := model.Profile{}

	if match != nil {
		age, err := strconv.Atoi(extractString(contents, ageCompile))
		if err == nil {
			profile.Age = age
		}
	}

	profile.Marriage = extractString(contents, nicknameCompile)

	result := engine.ParseResult{
		Items: []interface{}{profile},
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
