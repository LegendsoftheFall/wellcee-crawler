package parser

import (
	"bytes"
	"crawler/engine"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

const (
	DistUrl       = `https://www.wellcee.com/api/house/filter`
	MaxPage int64 = 1
)

// GetCityID 获取当前城市的ID
func GetCityID(u string) (id string) {
	rawUrl, _ := url.Parse(u)
	str := strings.Split(rawUrl.RawQuery, "&")
	id = strings.Split(str[0], "=")[1]
	return
}

func GetApartmentIDList(cityID string, page int64) (list []string, err error) {
	// 模拟发送请求
	var i int64
	for i = 1; i <= page; i++ {
		req := engine.ListRequest{
			CityID: cityID,
			Lang:   1,
			Pn:     page,
		}
		// 请求转为JSON格式
		reqJson, err1 := json.Marshal(req)
		if err1 != nil {
			log.Fatalf("json.Marshal err : %s", err1)
		}
		// 发送请求
		res, err2 := http.Post(DistUrl, "application/json", bytes.NewBuffer(reqJson))
		if err2 != nil {
			log.Fatalf("post %s err : %s", DistUrl, err2)
		}
		// 接收请求
		body, err3 := io.ReadAll(res.Body)
		if err3 != nil {
			log.Fatalf("io.ReadAll err : %s", err3)
		}

		// 写入结构体
		responseList := new(engine.ListResponse)
		_ = json.Unmarshal(body, responseList)

		for _, item := range responseList.Data.List {
			list = append(list, item.ID)
		}
	}
	fmt.Println(list)
	return
}

func ParseApartmentList(doc *html.Node, cityUrl string) (r engine.ParseResult) {
	// 解析出当前城市id
	cityID := GetCityID(cityUrl)
	// 获取第i页的apartmentID列表

	aIDs, err := GetApartmentIDList(cityID, MaxPage)
	if err != nil {
		log.Fatalf("GetApartmentIDList err : %s", err)
	}
	r = engine.ParseResult{
		Requests: nil,
		Items:    nil,
	}
	for _, id := range aIDs {
		url := fmt.Sprintf("https://www.wellcee.com/rent-apartment/%s?lang=zh", id)
		r.Requests = append(r.Requests, engine.Request{
			Url: url,
			ParseFunc: func(node *html.Node) engine.ParseResult {
				return ParseApartmentInfo(node, id, url)
			},
		})
	}

	return
}
