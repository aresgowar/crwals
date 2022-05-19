package dbconnect

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBConnect(name string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@/crawldata?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database") // Kiểm tra kết nối tới database
	}
	return db, err
}
