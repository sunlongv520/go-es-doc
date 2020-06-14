package Funs

import (
	"es.jtthink.com/AppInit"
	"es.jtthink.com/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
)

func SearchBook(ctx *gin.Context)  {
    searchModel:=Models.NewSearchModel()
    err:=ctx.BindJSON(searchModel)
	if err!=nil{
		log.Println(searchModel)
		log.Println(err)
		ctx.JSON(400,gin.H{"error":err})
		return
	}
    qList:=make([]elastic.Query,0)
    if searchModel.BookName !="" { //图书名
		machQuery:=elastic.NewMatchQuery("BookName",searchModel.BookName)
		qList=append(qList,machQuery)
	}
    if searchModel.BookPress!="" { //判断出版社
    	pressQuery:=elastic.NewTermQuery("BookPress",searchModel.BookPress)
		qList=append(qList,pressQuery)
	}
    if searchModel.BookPrice1Start>0 || searchModel.BookPrice1End>0{
    	priceRangeQuery:=elastic.NewRangeQuery("BookPrice1")
    	if searchModel.BookPrice1Start>0{
			priceRangeQuery.Gte(searchModel.BookPrice1Start)
		}
		if searchModel.BookPrice1End>0{
			priceRangeQuery.Lte(searchModel.BookPrice1End)
		}
		qList=append(qList,priceRangeQuery)
	}


	 boolMustQuery:=elastic.NewBoolQuery().Must(qList...)

	rsp,err:=AppInit.GetEsClient().Search().Query(boolMustQuery).
		Index("books").Do(ctx)
	log.Println(err)
	if err!=nil{
		fmt.Println("fewfwefwef",err)
		ctx.JSON(500,gin.H{"error":err})
	}else{
		ctx.JSON(200,gin.H{"result":MapToBooks(rsp)})
	}
}
