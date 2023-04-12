package engine

import (
	"crawler/model"

	"golang.org/x/net/html"
)

// Request 请求体：包含url和解析该url的解析方法 解析方法返回解析结果
type Request struct {
	Url       string
	ParseFunc func(*html.Node) ParseResult
}

// ParseResult 解析结果： 包含解析出的多个请求和解析出的目标类型
type ParseResult struct {
	Requests []Request
	Items    []model.ApartmentDetail
}

func NilParser(*html.Node) (pr ParseResult) {
	return
}

type ListResponse struct {
	Code int64 `json:"code"`
	Data struct {
		Count int64 `json:"count"`
		List  []struct {
			ID string `json:"id"`
		} `json:"list"`
	} `json:"data"`
	Msg string `json:"msg"`
	Ret bool   `json:"ret"`
}

type ListRequest struct {
	CityID string `json:"cityId"`
	Lang   int64  `json:"lang"`
	Pn     int64  `json:"pn"`
}
