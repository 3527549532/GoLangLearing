package main

import (
	"_course_pachong/controllers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 部署静态html资源
	controllers.R.Static("/static", "./static")
	
	controllers.ReceiveItPostRequest()
	controllers.ReceiveBookPostRequest()
	controllers.ReceiveKnowledgePostRequest()
	err := controllers.R.Run(":8080")
	if err != nil {
		return
	}

	// 爬虫获取数据
	//controllers.RequestItJsonDataBypachong()
	//controllers.RequestBookJsonDataBypachong()
	//controllers.RequestKlJsonDataBypachong()

}
