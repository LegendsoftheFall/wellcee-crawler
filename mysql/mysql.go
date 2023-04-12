package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

const (
	USERNAME    = "root"
	PASSWORD    = "20011112"
	HOST        = "8.134.222.37"
	PORT        = 3306
	DBNAME      = "wellcee"
	MaxConn     = 1000
	MaxIdleConn = 50
)

func InitDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		USERNAME,
		PASSWORD,
		HOST,
		PORT,
		DBNAME)
	// 也可以使用MustConnect连接不成功就panic
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("connect DB failed %s", err)
		return
	}
	DB.SetMaxOpenConns(MaxConn)
	DB.SetMaxIdleConns(MaxIdleConn)
	return
}

func Close() {
	_ = DB.Close()
}
