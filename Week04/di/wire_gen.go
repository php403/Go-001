// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/php403/go-001/week04/dao"
	"github.com/php403/go-001/week04/redis"
)


func InitApp() (*WireStruct, error) {
	options := NewAppOption()
	daoOptions := dao.NewOptionsMysql()
	redisOptions := redis.NewOptionsRedis()
	wireStruct := &WireStruct{
		App:   options,
		Mysql: daoOptions,
		Redis: redisOptions,
	}
	return wireStruct, nil
}

// wire.go:

type WireStruct struct {
	App   *Options
	Mysql *dao.Options
	Redis *redis.Options
}
