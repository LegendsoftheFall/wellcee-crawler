package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	// 队列
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 队列中有请求就解析
	for len(requests) > 0 {
		// 取出第一个请求
		r := requests[0]
		requests = requests[1:]
		// 交给Worker解析
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		// Parser解析出新的url则加入请求队列
		requests = append(requests, parseResult.Requests...)

		// 打印已解析的目标
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	// 交给fetcher爬取请求中url的内容
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Fatalf("Fetcher :Error"+"fetching url %s:%v", r.Url, err)
		return ParseResult{}, err
	}

	// 无误则交给Parser
	return r.ParseFunc(body), nil
}
