package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/opentracing/opentracing-go"
	"github.com/siliconvalley001/wen/web/router"
	"github.com/siliconvalley001/wen/web/common"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/registry"
	"log"


)

func main() {
	//配置中心
	//consulconfig,err:=common.GetConsulConfig("127.0.0.1",8500,"/lizhengda/project/config")
	//if err!=nil{
	//	//log.Fatal(err)
	//}
	//注册中心
	registry:=consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs=[]string{
				"127.0.0.1:8500",
			}

		},
	)
	//链路追踪
	t,io,err:=common.NewTrace("micro.product.jaeger","127.0.0.1:6831")
	if err!=nil{
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	r := gin.Default()

	router.InitRouter(r)
	// Create service
	srv := web.NewService(
		web.Handler(r),
		web.Name("web.micro"),
		web.Version("laster"),
		web.Address(":8089"),
		web.Registry(registry),
		//web.MicroService()
	)

	if err:=srv.Init();err!=nil{
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
