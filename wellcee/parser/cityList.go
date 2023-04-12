package parser

import (
	"crawler/engine"
	"fmt"
	"log"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const (
	CityQuery = `//div[@class="footer-city J_footer_find"]/div[@class="col"]/p/a`
)

func ParseCityList(doc *html.Node) (r engine.ParseResult) {
	r = engine.ParseResult{
		Requests: nil,
		Items:    nil,
	}
	nodes, err := htmlquery.QueryAll(doc, CityQuery)
	if err != nil {
		log.Fatalf("ParseCityList err : %s\n", err)
	}
	for _, n := range nodes {
		name := htmlquery.InnerText(n)
		url := fmt.Sprintf("https://www.wellcee.com%s&lang=zh", htmlquery.SelectAttr(n, "href"))
		log.Printf("city: %s, name:%s \n", url, name)
		r.Requests = append(r.Requests, engine.Request{
			Url: url,
			ParseFunc: func(node *html.Node) engine.ParseResult {
				return ParseApartmentList(node, url)
			},
		})
	}
	return
}
