package main

import (
	"github.com/tang-go/go-dog/log"
	"github.com/tang-go/go-dog/pkg/client"
	"github.com/tang-go/go-dog/pkg/config"
	"github.com/tang-go/go-dog/pkg/discovery"
	"github.com/tang-go/go-dog/pkg/register"
	"github.com/tang-go/go-dog/pkg/service"
	"github.com/tang-go/go-dog/plugins"
)

const (
	//AddErr 相加失败
	AddErr = 10001
)

//AddReq 管理员登录
type AddReq struct {
	A int `json:"a" description:"数字A" type:"int"`
	B int `json:"b" description:"数字B" type:"int"`
}

//AddRes 管理员登录返回
type AddRes struct {
	Result int `json:"result" description:"结果" type:"int"`
}

//Add 相加
func Add(ctx plugins.Context, request AddReq) (response AddRes, err error) {
	//错误返回
	//err = customerror.EnCodeError(define.AddErr, "相加失败")
	log.Traceln("收到请求", request)
	response.Result = request.A + request.B
	return
}

//使用etcd为服务注册中心
func main() {
	//初始化配置
	cfg := config.NewConfig()
	//etcd 启动etcd服务端
	ser := service.CreateService(
		//名称
		"etcd",
		//配置
		cfg,
		//客户端
		client.NewClient(
			//配置
			cfg,
			//使用etcd服务发现
			discovery.NewEtcdDiscovery(cfg.GetDiscovery()),
		),
		//etcd服务注册插件
		register.NewEtcdRegister(cfg.GetDiscovery()))
	//注册RPC函数
	ser.RPC(
		//方法名称
		"Add",
		//等级
		3,
		//是否需要验证
		false,
		//说明
		"相加",
		//方法
		Add,
	)
	err := ser.Run()
	if err != nil {
		log.Traceln("服务退出", err.Error())
	}
}
