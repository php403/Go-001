package di

import (
	"github.com/google/wire"
	"github.com/php403/go-001/week04/dao"
	"github.com/php403/go-001/week04/redis"
)



/*grpc.ProviderSet,
app.ProviderSet,*/

func InitApp() (func(), error) {
	panic(wire.Build(dao.ProviderSet,redis.ProviderSet))
}