package main

import (
        "github.com/micro/go-log"
        "github.com/micro/go-web"
        "github.com/julienschmidt/httprouter"
        "sss/IhomeWeb/handler"
        "net/http"
        _"sss/IhomeWeb/models"
)

func main() {
	// create new web service
	//创建服务
        service := web.NewService(
                web.Name("go.micro.web.IhomeWeb"),
                web.Version("latest"),
                web.Address(":10010"),
        )

	// initialise service
	//服务初始化
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	//// register html handler
	//映射静态文件
	//service.Handle("/", http.FileServer(http.Dir("html")))
	//注册路由
	//// register call handler
	//service.HandleFunc("/example/call", handler.ExampleCall)

	//创建句柄
	rou:=httprouter.New()
	///映射静态文件
	rou.NotFound=http.FileServer(http.Dir("html"))
	//注册路由
	rou.GET("/example/call",handler.ExampleCall)


	//请求地域信息
	rou.GET("/api/v1.0/areas",handler.GetArea)

	//请求首页轮播图
	rou.GET("/api/v1.0/house/index",handler.GetIndex)

	//请求session信息
	rou.GET("/api/v1.0/session",handler.GetSession)

	//请求验证码图片
	rou.GET("/api/v1.0/imagecode/:uuid",handler.GetImageCd)

	//请求短信服务
	rou.GET("/api/v1.0/smscode/:mobile",handler.GetSmscd)

	//发送注册信息服务
	rou.POST("/api/v1.0/users",handler.PostRet)

	service.Handle("/",rou)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
