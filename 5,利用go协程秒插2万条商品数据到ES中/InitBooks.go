package main

import (
	"context"
	"es.jtthink.com/AppInit"
	"es.jtthink.com/Models"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"strconv"
	"sync"
)

func main()  {

	page:=1
	pagesize:=500
	wg:=sync.WaitGroup{}
	for{
		book_list:=Models.BookList{}
		db:=AppInit.GetDB().Order("book_id desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&book_list)
		if db.Error!=nil || len(book_list)==0{
			break
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			client:=AppInit.GetEsClient()
			bulk:=client.Bulk()
			for _,book:=range book_list {
				req:=elastic.NewBulkIndexRequest()
				req.Index("books").Id(strconv.Itoa(book.BookID)).Doc(book)
				bulk.Add(req)
			}
			rsp,err:=bulk.Do(context.Background())
			if err!=nil {
				log.Println(err)
			}else {
				fmt.Println(rsp)
			}
		}()
       page=page+1  //必须有
	}
		wg.Wait()

}
