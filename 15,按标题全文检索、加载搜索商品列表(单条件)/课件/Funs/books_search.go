package Funs

import (
	"es.jtthink.com/AppInit"
	"es.jtthink.com/Models"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

func SearchBook(ctx *gin.Context)  {
    searchModel:=Models.NewSearchModel()
    err:=ctx.BindJSON(searchModel)
	if err!=nil{
		ctx.JSON(500,gin.H{"error":err})
		return
	}
	machQuery:=elastic.NewMatchQuery("BookName",searchModel.BookName)
	rsp,err:=AppInit.GetEsClient().Search().Query(machQuery).
		Index("books").Do(ctx)
	if err!=nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":MapToBooks(rsp)})
	}
}
