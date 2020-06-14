package main

import (
	"es.jtthink.com/Funs"
	"github.com/gin-gonic/gin"
)

func main()  {

	router:=gin.Default()
	g:=router.Group("/books")
	{
		g.Handle("GET","",Funs.LoadBook)
	}
	router.Run(":8080")


}
