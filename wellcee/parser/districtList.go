package parser

import (
	"crawler/engine"
	"fmt"

	"github.com/antchfx/htmlquery"

	"golang.org/x/net/html"
)

// https://www.wellcee.com/rent-apartment/shenzhen/list?cityId=15397547907923190&city=Shenzhen-apartment-house-rental-price-list#districtIds=15397548274301341
// https://www.wellcee.com/rent-apartment/beijing/list?cityId=15102232786514309&city=Beijing-apartment-house-rental-price-list&lang=zh#districtIds=15133093880781510

//<a href="/rent-apartment/shanghai/list?cityId=15102233103895305">上海</a>

const (
	DistrictQuery = `//div[@class="strategy-content-box"]//p/a`
)

func ParseDistrictList(doc *html.Node) (r engine.ParseResult) {
	nodes, err := htmlquery.QueryAll(doc, DistrictQuery)
	if err != nil {
		panic(err)
	}
	for _, node := range nodes[:len(nodes)-1] {
		fmt.Println(htmlquery.InnerText(node), htmlquery.SelectAttr(node, "href"))
	}
	return
}
