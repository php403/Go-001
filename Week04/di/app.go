package di

import (
	"context"
	"github.com/php403/go-001/week04/dao"
	"github.com/php403/go-001/week04/service"
	"github.com/php403/go-001/week04/transport"
	"github.com/php403/go-001/week04/endpoint"
	"time"
)


type App struct {

}

func NewApp() (closeFunc func(), err error){
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
	return func() {

	},nil

}
