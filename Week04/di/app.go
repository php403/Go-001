package di

import (
	"context"
	"flag"
	"github.com/google/wire"
	"github.com/php403/go-001/week04/dao"
	"github.com/php403/go-001/week04/service"
	"github.com/php403/go-001/week04/transport"
	"github.com/php403/go-001/week04/endpoint"
	"net/http"
	"strconv"
	"time"
)

var ProviderSet = wire.NewSet(NewApp,NewAppOption)



type Options struct {
	Port int
}

func NewAppOption() *Options {
	return &Options{10086}
}

func NewApp(o *Options) (closeFunc func(), err error){
	var (
		// 服务地址和服务名
		servicePort = flag.Int("service.port", o.Port, "service port")
	)

	ctx := context.Background()

	userService := service.MakeUserServiceImpl(&dao.UserDAOImpl{})

	userEndpoints := &endpoint.UserEndpoints{
		endpoint.MakeRegisterEndpoint(userService),
		endpoint.MakeLoginEndpoint(userService),
	}
	closeFunc = func() {
		_, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		cancel()
	}
	transport.MakeHttpHandler(ctx, userEndpoints)
	r := transport.MakeHttpHandler(ctx, userEndpoints)
	errChan := make(chan error)
	go func() {
		errChan <- http.ListenAndServe(":"  + strconv.Itoa(*servicePort), r)
	}()
	return func() {

	},nil

}
