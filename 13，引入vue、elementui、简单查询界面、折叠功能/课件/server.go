package main

import (
	"es.jtthink.com/Funs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {


	router:=gin.Default()
	g:=router.Group("/books") //业务相关
	{
		//图书列表
		g.Handle("GET","",Funs.LoadBook)
		g.Handle("GET","/press/:press",Funs.LoadBookByPress)
	}
	//helper
	helper:=router.Group("/helper")
	{
		helper.Handle("GET","/press",Funs.PressList)

	}
	router.StaticFS("/ui", http.Dir("./htmls"))


	router.Run(":8080")


}
