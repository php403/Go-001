// +build wireinject
package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var ProviderSet = wire.NewSet(NewMysql, NewOptionsMysql)
var db *gorm.DB

// Options is  configuration of database
type Options struct {
	Host string
	Port int
	User string
	PassWd string
	DbName string
}

func NewOptionsMysql() *Options {
	return &Options{"139.9.141.27", 3306, "pg719", "pg719@1996", "demo"}
}

// Init 初始化数据库
func NewMysql(o *Options) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", o))
	if err != nil{
		log.Println(err)
		return db,err
	}
	db.SingularTable(true)
	return db,err
}


