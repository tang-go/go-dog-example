package main

import (
	"github.com/tang-go/go-dog/log"
	"github.com/tang-go/go-dog/mysql"
	"github.com/tang-go/go-dog/pkg/service"
	"github.com/tang-go/go-dog/plugins"
)

const (
	//AddErr 相加失败
	AddErr = 10001
)

var gMysql *mysql.Mysql

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

func main() {
	//mysql 为启动服务的名称
	ser := service.CreateService("mysql")
	//注册RPC函数
	ser.RPC(
		"Add", //方法名称
		3,     //等级
		false, //是否需要验证
		"相加",  //说明
		Add)   //方法
	//初始化数据库
	gMysql = mysql.NewMysql(ser.GetCfg())
	//获取读数据数据库对象
	gMysql.GetReadEngine()
	//获取写数据库对象
	gMysql.GetWriteEngine()
	err := ser.Run()
	if err != nil {
		log.Traceln("服务退出", err.Error())
	}
}
