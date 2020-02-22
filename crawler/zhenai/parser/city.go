package parser

import (
	"eg/egolang/crawler/engine"
	"regexp"
)

var (
	cityReCompiler    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	cityUrlReCompiler = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, url string) engine.ParseResult {
	matches := cityReCompiler.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		// result.Items = append(result.Items, "User: "+name)
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ProfileParser(name),
			},
		)
	}

	matches = cityUrlReCompiler.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
