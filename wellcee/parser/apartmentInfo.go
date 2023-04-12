package parser

import (
	"crawler/engine"
	"crawler/model"
	"log"

	"github.com/antchfx/htmlquery"

	"golang.org/x/net/html"
)

const (
	ImageQuery    = `//div[@class="clearfix photo-box9"]/a[@rel="photo-list"]`
	VideoQuery    = `//video/source[@type="video/mp4"]`
	DetailQuery   = `//div[@class="detail"]`
	TagsQuery     = `//strong[@class="tag"]`
	ListQuery     = `//div[@class="list clearfix"]/div[@class="item"]/span[@class="value"]`
	IntroQuery    = `//pre[@class="house-intro"]`
	PriceQuery    = `//p[@class="price"]`
	StatusQuery   = `//div[@class="detail-status"]`
	LocationQuery = `//p[@class="house-address"]`
)

func ParseApartmentInfo(doc *html.Node, id, url string) (r engine.ParseResult) {
	detail := model.ApartmentDetail{}
	detail.ID = id
	detail.Url = url
	// 解析图片
	//nodes, err := htmlquery.QueryAll(doc, ImageQuery)
	//if err != nil {
	//	log.Fatalf("ParseHouseInfo err : %s\n", err)
	//}
	//for _, node := range nodes {
	//	//fmt.Println(htmlquery.SelectAttr(node, "href"))
	//	detail.Image = append(detail.Image, htmlquery.SelectAttr(node, "href"))
	//}
	// 解析视频
	//nodes, err = htmlquery.QueryAll(doc, VideoQuery)
	//if err != nil {
	//	log.Fatalf("ParseHouseInfo err : %s\n", err)
	//}
	//for _, node := range nodes {
	//	//fmt.Println(htmlquery.SelectAttr(node, "src"))
	//	detail.Video = append(detail.Video, htmlquery.SelectAttr(node, "src"))
	//}
	// 解析详情
	nodes, err := htmlquery.QueryAll(doc, DetailQuery)
	if err != nil {
		log.Fatalf("ParseHouseInfo err : %s\n", err)
	}
	for _, node := range nodes {
		// 出租状态
		statusHtml := htmlquery.FindOne(node, StatusQuery)
		detail.Status = htmlquery.InnerText(statusHtml)
		priceHtml := htmlquery.FindOne(node, PriceQuery)
		detail.Price = htmlquery.InnerText(priceHtml)
		introHtml := htmlquery.FindOne(node, IntroQuery)
		detail.Intro = htmlquery.InnerText(introHtml)
		locationHtml := htmlquery.FindOne(node, LocationQuery)
		detail.Location = htmlquery.InnerText(locationHtml)
	}
	// 解析标签
	//nodes, err = htmlquery.QueryAll(doc, TagsQuery)
	//if err != nil {
	//	log.Fatalf("ParseHouseInfo err : %s\n", err)
	//}
	//for _, node := range nodes {
	//	detail.Tag = append(detail.Tag, htmlquery.InnerText(node))
	//}

	// 解析列表
	nodes, err = htmlquery.QueryAll(doc, ListQuery)
	if err != nil {
		log.Fatalf("ParseHouseInfo err : %s\n", err)
	}
	for i, node := range nodes {
		switch i {
		case 0:
			detail.Type = htmlquery.InnerText(node)
		case 1:
			detail.Deposit = htmlquery.InnerText(node)
		case 2:
			detail.Room = htmlquery.InnerText(node)
		case 3:
			detail.Area = htmlquery.InnerText(node)
		case 4:
			detail.Subway = htmlquery.InnerText(node)
		case 5:
			detail.Floor = htmlquery.InnerText(node)
		default:
			continue
		}
	}

	// 解析标签
	//nodes, err = htmlquery.QueryAll(doc, IntroQuery)
	//if err != nil {
	//	log.Fatalf("ParseHouseInfo err : %s\n", err)
	//}
	//for _, node := range nodes {
	//	detail.Intro = htmlquery.InnerText(node)
	//}
	r.Items = append(r.Items, detail)

	return
}
