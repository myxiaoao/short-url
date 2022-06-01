package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Url struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Short    string `json:"short" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/short_url?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Url{})
	if err != nil {
		fmt.Println(err)
	}
}
