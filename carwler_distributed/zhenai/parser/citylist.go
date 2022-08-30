package parser

import (
	"pachong/carwler_distributed/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	matchs := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 2
	for _, match := range matchs {
		result.Item = append(result.Item, "city="+string(match[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}

	}

	return result
}
