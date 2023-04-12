package main

import (
	"crawler/engine"
	"crawler/mysql"
	"crawler/scheduler"
	"crawler/storage"
	"crawler/wellcee/parser"
	"fmt"
)

//Url:       "https://www.wellcee.com/rent-apartment/shanghai/list?cityId=15102233103895305&lang=zh",
//ParseFunc: parser.ParseHouseList,
//Url:       "https://www.wellcee.com/rent-apartment/1680005117598815?lang=zh",
//ParseFunc: parser.ParseHouseInfo,

func main() {
	// 连接数据库
	if err := mysql.InitDB(); err != nil {
		fmt.Printf("init mysql error, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 启动爬虫
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChannel: storage.ItemSaver(),
	}
	//e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:       "https://www.wellcee.com/shenzhen/rent-apartment/zh?lang=zh",
		ParseFunc: parser.ParseCityList,
	})
}
