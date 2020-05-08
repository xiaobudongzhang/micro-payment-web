package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/xiaobudongzhang/micro-basic/basic"
	"github.com/xiaobudongzhang/micro-basic/config"
	"github.com/xiaobudongzhang/micro-payment-web/handler"
)

func main() {
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.payment"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8090"),
	)

	// initialise service
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	authHandler := http.HandlerFunc(handler.PayOrder)
	service.Handle("/payment/pay-order", handler.AuthWrapper(authHandler))
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
