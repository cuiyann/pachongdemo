package engine

import (
	"log"
	"pachong/carwler_distributed/fetcher"
)

func Run(seed ...Request) {
	var requests []Request
	for _, request := range seed {
		requests = append(requests, request)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Println("url = " + r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("错误信息=%v,url=%s", err, r.Url)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Request...)
		for _, i := range parseResult.Item {
			log.Printf("获取item %v", i)
		}
	}

}
