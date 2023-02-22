package main

import (
	"chat/app/http/middleware"
	"chat/config"
	"chat/route"
	"fmt"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func init() {
	//初始化配置文件
	config.Setup()
}

func main() {

	gin.SetMode(config.Conf.HttpServer.RunModel)

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())

	route.APIRouter(r)
	route.WebRouter(r)

	readTimeout := config.Conf.HttpServer.ReadTimeout * time.Second
	writeTimeout := config.Conf.HttpServer.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", config.Conf.HttpServer.HttpPort)

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	server := endless.NewServer(endPoint, r)
	//在 BeforeBegin 时输出当前进程的 pid，调用 ListenAndServe 将实际“启动”服务
	server.BeforeBegin = func(add string) {
		fmt.Printf("actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("server start err:%v", err)
	} else {
		fmt.Println("server start success")
	}

}
