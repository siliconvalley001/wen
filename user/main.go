package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/opentracing/opentracing-go"
	_"github.com/siliconvalley001/project1/user/model"
	"github.com/siliconvalley001/project1/user/common"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/registry"

	handler1 "github.com/siliconvalley001/project1/user/handler"

	"fmt"
	ex "github.com/siliconvalley001/project1/user/proto"
)

func main() {
	//配置中心
	consulconfig,err:=common.GetConsulConfig("127.0.0.1",8500,"/lizhengda/project/config")
	if err!=nil{
		//log.Fatal(err)
	}
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
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)





	// Create service
	srv:=micro.NewService(
		micro.Name("micro.user"),
		micro.Version("latest"),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.Registry(registry),
		micro.Config(consulconfig),

	)
	srv.Init()
	// Register handler
	//pb.RegisterUserHandler(srv.Server(), new(handler.User))

	ex.RegisterUserHandler(srv.Server(),new(handler1.Handler))
	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
