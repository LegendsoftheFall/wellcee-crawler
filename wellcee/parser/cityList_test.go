package parser

import (
	"crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := fetcher.Fetch("https://www.wellcee.com/?lang=zh")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(content)
	const resultSize = 22
	expectedUrls := []string{
		"https://www.wellcee.com/shanghai/rent-apartment",
		"https://www.wellcee.com/beijing/rent-apartment",
		"https://www.wellcee.com/shenzhen/rent-apartment",
	}
	expectedCities := []string{
		"上海", "北京", "深圳",
	}

	for i, u := range expectedUrls {
		if result.Requests[i].Url != u {
			t.Errorf("expected url %d:%s but get %s", i, u, result.Requests[i].Url)
		}
	}
	for i, c := range expectedCities {
		if result.Items[i] != c {
			t.Errorf("expected city %d:%s but get %s", i, c, result.Items[i])
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d"+"requests; but had %d",
			resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d"+"items; but had %d",
			resultSize, len(result.Items))
	}
}
