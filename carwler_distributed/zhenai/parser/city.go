package parser

import (
	"pachong/carwler_distributed/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]* target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRe)
	matchs := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, match := range matchs {
		name := string(match[2])
		result.Item = append(result.Item, "user="+string(match[2]))
		result.Request = append(result.Request, engine.Request{
			Url: string(match[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	return result
}
