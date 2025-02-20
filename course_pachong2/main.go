package main

import (
	"_course_pachong/controllers"
)

func main() {
	controllers.ReceiveItPostRequest()
	//controllers.ReceiveBookPostRequest()

	//jsonData := `[
	//	{
	//		"title": "\u8bfe\u7a0b\u540d\u79f0",
	//		"slot": "show_name"
	//	},
	//	{
	//		"title": "\u7f51\u76d8\u5bc6\u7801",
	//		"slot": "pwd"
	//	},
	//	{
	//		"title": "\u4ef7\u683c",
	//		"key": "price"
	//	},
	//	{
	//		"title": "\u79ef\u5206\u5151\u6362",
	//		"key": "costscore"
	//	},
	//	{
	//		"title": "\u5185\u5bb9\u5927\u5c0f",
	//		"key": "size"
	//	},
	//	{
	//		"title": "\u6807\u7b7e",
	//		"key": "tag"
	//	},
	//	{
	//		"title": "\u64cd\u4f5c",
	//		"slot": "action"
	//	}
	//]`
	//
	//// 解析JSON数据到结构体切片
	//var configs []ResourceColumnConfig
	//err := json.Unmarshal([]byte(jsonData), &configs)
	//if err != nil {
	//	fmt.Println("Error unmarshalling JSON:", err)
	//	return
	//}
	//
	//dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// 自动迁移模式（会根据当前的Go结构体自动创建或更新表）
	//err = db.AutoMigrate(&ResourceColumnConfig{})
	//if err != nil {
	//	fmt.Println("Error migrating database:", err)
	//	return
	//}
	//
	//// 将结构体切片保存到数据库
	//result := db.Create(&configs)
	//if result.Error != nil {
	//	fmt.Println("Error creating records in database:", result.Error)
	//	return
	//}

	//部署静态html资源
	//r := gin.Default()
	//// 将 / 映射到 static 目录，这样可以直接访问 static 下的所有文件
	//r.Static("/static", "./static")
	//
	////// 启动服务器
	//r.Run(":8080")
	////部署资源配置结束
	//
	//url := "https://dl.zzyyww.cn/learn/api/itcourselist/"
	//method := "POST"
	//
	//payload := strings.NewReader(`{"learn_type":"type_it","key1_1":"","key1_2":"","key1_3":"","it_current":1,"sel_ittags":[]}`)
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//req.Header.Add("Accept", "application/json, text/plain, */*")
	//req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	//req.Header.Add("Origin", "https://dl.zzyyww.cn")
	//req.Header.Add("Referer", "https://dl.zzyyww.cn/learn/")
	//req.Header.Add("Sec-Fetch-Dest", "empty")
	//req.Header.Add("Sec-Fetch-Mode", "cors")
	//req.Header.Add("Sec-Fetch-Site", "same-origin")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	//req.Header.Add("sec-ch-ua", "\"Not(A:Brand\";v=\"99\", \"Google Chrome\";v=\"133\", \"Chromium\";v=\"133\"")
	//req.Header.Add("sec-ch-ua-mobile", "?0")
	//req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	////
	//
	//// 解析JSON数据
	//var courseList ItCourseList
	//err = json.Unmarshal(body, &courseList)
	//if err != nil {
	//	log.Fatalf("解析JSON数据失败: %v", err)
	//}
	//
	////// 打印解析后的数据
	////fmt.Printf("接口返回的数据: %+v\n", courseList)
	//// 数据库连接信息
	//dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database"
	//db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//// 检查数据库连接
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 遍历 Itresourcedata 数组并插入到数据库
	//for _, resource := range courseList.Itresourcedata {
	//	// 准备插入SQL语句（使用参数化查询）
	//	stmt, err := db.Prepare("INSERT INTO itresourcedata (lrid, show_name, share_url, pwd, price, costscore, size, tag) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer stmt.Close()
	//
	//	// 执行插入操作
	//	_, err = stmt.Exec(resource.Lrid, resource.ShowName, resource.ShareURL, resource.Pwd, resource.Price, resource.Costscore, resource.Size, resource.Tag)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//
	//fmt.Println("ItResourceData array saved successfully")
}
