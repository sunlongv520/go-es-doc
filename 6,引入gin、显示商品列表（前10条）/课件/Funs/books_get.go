package Funs

import (
	"es.jtthink.com/AppInit"
	"github.com/gin-gonic/gin"
)

func LoadBook(ctx *gin.Context)  {

	rsp,err:=AppInit.GetEsClient().Search().Index("books").Do(ctx)
	if err!=nil{
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":rsp.Hits.Hits})
	}
}