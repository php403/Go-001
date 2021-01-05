// +build wireinject

package di
import (
	"github.com/google/wire"
	"github.com/php403/go-001/week04/dao"
	"github.com/php403/go-001/week04/redis"
)

type WireStruct struct {
	App *Options
	Mysql *dao.Options
	Redis *redis.Options
}

func InitApp() (*WireStruct,error) {
	panic(wire.Build(ProviderSet,dao.ProviderSet,redis.ProviderSet,WireStruct{}))
}