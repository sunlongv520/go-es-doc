package main

import (
	"context"
	"es.jtthink.com/AppInit"
	"es.jtthink.com/Models"
	"fmt"
	"strconv"
)

func main()  {


	page:=1
	pagesize:=50
	for{
		book_list:=Models.BookList{}
		db:=AppInit.GetDB().Order("book_id desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&book_list)
		if db.Error!=nil || len(book_list)==0{
			break
		}
		for _,book:=range book_list {
			ctx:=context.Background()
			rsp,err:=AppInit.GetEsClient().Index().Index("books").
				Id(strconv.Itoa(book.BookID)).BodyJson(book).Do(ctx)
			if err!=nil{
				fmt.Println(err)
			}else {
				fmt.Println(rsp)
			}
		}



		break
	}
}
