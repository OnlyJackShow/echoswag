package main

import (
	_ "echoswag/docs"
	"echoswag/route"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"net/url"
)


func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//错误中间件
	e.Use(middleware.Recover())
	surl, err := url.Parse("http://10.1.1.21:8090")
	if err != nil {
		fmt.Println(err.Error())
	}

	//代理
	e.Group("/first", middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
		{URL: surl},
	})))



	//路由
	e = route.Route(e)
	e.Logger.Fatal(e.Start(":1323"))
}

