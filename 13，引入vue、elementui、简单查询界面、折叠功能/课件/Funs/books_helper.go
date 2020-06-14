package Funs

import (
	"es.jtthink.com/AppInit"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

//无脑取前10个出版社
func PressList(ctx *gin.Context)  {
	cb:=elastic.NewCollapseBuilder("BookPress")
	rsp,err:=AppInit.GetEsClient().Search().Collapse(cb).FetchSource(false).Index("books").Do(ctx)
	if err!=nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":rsp})
	}
}
