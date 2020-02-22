package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Id      string
	Type    string
	Url     string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
