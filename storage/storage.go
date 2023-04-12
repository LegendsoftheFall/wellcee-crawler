package storage

import (
	"crawler/model"
	"crawler/mysql"
	"log"
	"strconv"
)

func ItemSaver() chan model.ApartmentDetail {
	in := make(chan model.ApartmentDetail)
	go func() {
		itemCount := 0
		for {
			item := <-in
			log.Printf("Storage: Got item #%d %v", itemCount, item)
			itemCount++
			if err := InsertMysql(item); err != nil {
				log.Fatalf("Save error: item %v,err %v", item, err)
			}
		}
	}()
	return in
}

func InsertMysql(detail model.ApartmentDetail) (err error) {
	id, _ := strconv.Atoi(detail.ID)
	sqlStr := `insert into apartment  (aid,url,type,floor,location,deposit,room,
                              area,subway,status,price,intro) values (?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err = mysql.DB.Exec(sqlStr, id, detail.Url, detail.Type, detail.Floor, detail.Location, detail.Deposit,
		detail.Room, detail.Area, detail.Subway, detail.Status, detail.Price, detail.Intro)
	return
}
