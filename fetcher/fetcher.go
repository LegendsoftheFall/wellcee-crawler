package fetcher

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/antchfx/htmlquery"

	"golang.org/x/net/html"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) (*html.Node, error) {
	//<-rateLimiter
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil) //构造GET请求，获取信息
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36 Edg/104.0.1293.47")

	//resp, err := http.Get(url)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return htmlquery.Parse(resp.Body)
}
