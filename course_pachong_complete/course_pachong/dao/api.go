package dao

import (
	"_course_pachong/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func SearchIt(jsonRequestItData models.RequestITBodyData) []byte {

	// 连接到MySQL数据库
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 查询所有ResourceColumnConfig记录
	var configs []models.ItResourceDataColumn
	jsonData := `[
        {
            "title": "\u8bfe\u7a0b\u540d\u79f0",
            "slot": "show_name"
        },
        {
            "title": "\u7f51\u76d8\u5bc6\u7801",
            "slot": "pwd"
        },
        {
            "title": "\u4ef7\u683c",
            "key": "price"
        },
        {
            "title": "\u79ef\u5206\u5151\u6362",
            "key": "costscore"
        },
        {
            "title": "\u5185\u5bb9\u5927\u5c0f",
            "key": "size"
        },
        {
            "title": "\u6807\u7b7e",
            "key": "tag"
        },
        {
            "title": "\u64cd\u4f5c",
            "slot": "action"
        }
    ]`

	// 解析 It 标签列 Json 数据
	err1 := json.Unmarshal([]byte(jsonData), &configs)
	if err1 != nil {
		fmt.Println("Error parsing JSON:", err1)
	}

	//查询所有ItResourceData记录
	var itresourcedata []models.ItResourceData
	itresourcedataResult := db.Offset((jsonRequestItData.ITCurrent - 1) * 25).Limit(25).Find(&itresourcedata)
	if itresourcedataResult.Error != nil {
		log.Fatalf("failed to query database: %v", itresourcedataResult.Error)
	}

	// 定义一个变量来存储条目数
	var totalCount int64

	// 使用 Count 方法获取条目数
	result := db.Model(&models.ItResourceData{}).Count(&totalCount).Error
	if result != nil {
		log.Fatalf("failed to query database: %v", result)
	}

	response := models.ResponseItData{
		ItResourceColNums: configs,
		ItResourceData:    itresourcedata,
		ItResourceCount:   totalCount,
	}

	// 序列化为JSON格式
	jsonData1, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal JSON: %v", err)
	}
	return jsonData1
}

func SearchBook(jsonRequestData models.RequestBookBodyData) []byte {
	// 连接到MySQL数据库
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 定义一个变量来存储条目数
	var totalCount int64

	// 使用 Count 方法获取条目数
	result := db.Model(&models.BookResourceData{}).Count(&totalCount).Error
	if result != nil {
		log.Fatalf("failed to connect database: %v", result)
	}

	var BookResourceDataColumn []models.BookResourceDataColumn
	jsonData := `[
		{
			"title": "\u8bfe\u7a0b\u540d\u79f0",
			"slot": "show_name"
		},
		{
			"title": "\u7f51\u76d8\u5bc6\u7801",
			"slot": "pwd"
		}
	]`
	// 解析JSON数据
	err1 := json.Unmarshal([]byte(jsonData), &BookResourceDataColumn)
	if err1 != nil {
		fmt.Println("Error parsing JSON:", err1)
	}

	//查询所有BookResourceData记录
	var BookResourceData []models.BookResourceData
	BookResourceDataResult := db.Offset((jsonRequestData.BookCurrent - 1) * 25).Limit(25).Find(&BookResourceData)
	if BookResourceDataResult.Error != nil {
		log.Fatalf("failed to query database: %v", BookResourceDataResult.Error)
	}

	response := models.BookResourceDataResponseData{
		BookResourceColNums: BookResourceDataColumn,
		BookResourceData:    BookResourceData,
		BookResourceCount:   totalCount,
	}

	// 序列化为JSON格式
	jsonData1, err2 := json.MarshalIndent(response, "", "    ")
	if err2 != nil {
		log.Fatalf("failed to marshal JSON: %v", err2)
	}

	return jsonData1
}

func SearchKnowledge(jsonRequestData models.RequestKnowledgeData) []byte {
	// 连接到MySQL数据库
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	var KnowledgeResourceDataColumn []models.KnowledgeResourceDataColumn
	jsonData := `[
        {
            "title": "\u8bfe\u7a0b\u540d\u79f0",
            "slot": "show_name"
        },
        {
            "title": "\u7f51\u76d8\u5bc6\u7801",
            "slot": "pwd"
        },
        {
            "title": "\u4ef7\u683c",
            "key": "price"
        },
        {
            "title": "\u79ef\u5206\u5151\u6362",
            "key": "costscore"
        },
        {
            "title": "\u5185\u5bb9\u5927\u5c0f",
            "key": "size"
        },
        {
            "title": "\u6807\u7b7e",
            "key": "tag"
        },
        {
            "title": "\u64cd\u4f5c",
            "slot": "action"
        }
    ]`

	// 解析JSON数据
	err1 := json.Unmarshal([]byte(jsonData), &KnowledgeResourceDataColumn)
	if err1 != nil {
		fmt.Println("Error parsing JSON:", err1)
	}

	// 定义一个变量来存储条目数
	var totalCount int64

	// 使用 Count 方法获取条目数
	result := db.Model(&models.KnowledgeResourcesData{}).Count(&totalCount).Error
	if result != nil {
		panic(result) // 处理错误
	}

	//查询所有BookResourceData记录
	var KnowledgeResourceData []models.KnowledgeResourcesData
	KnowledgeResourceDataResult := db.Offset((jsonRequestData.KlCurrent - 1) * 25).Limit(25).Find(&KnowledgeResourceData)
	if KnowledgeResourceDataResult.Error != nil {
		log.Fatalf("failed to query database: %v", KnowledgeResourceDataResult.Error)
	}

	response := models.KnowledgeResourceDataResponseData{
		KnowledgeResourceColNums: KnowledgeResourceDataColumn,
		KnowledgeResourceData:    KnowledgeResourceData,
		KnowledgeResourceCount:   totalCount,
	}

	// 序列化为JSON格式
	jsonData1, err2 := json.MarshalIndent(response, "", "    ")
	if err2 != nil {
		log.Fatalf("failed to marshal JSON: %v", err2)
	}
	return jsonData1
}

func InserKlData(klData []models.KnowledgeResourcesData) {

	// 数据库连接
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 KnowledgeResourcesData 数组并插入到数据库
	for _, resource := range klData {
		// 准备插入SQL语句（使用参数化查询）
		stmt, err := db.Prepare("INSERT INTO knowledge_resources_data (lrid, show_name, share_url, pwd, price, costscore, size, tag) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		// 执行插入操作
		_, err = stmt.Exec(resource.Lrid, resource.ShowName, resource.ShareURL, resource.Pwd, resource.Price, resource.CostScore, resource.Size, resource.Tag)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("KnowledgeResourcesData array saved successfully")
}

func InserItData(itData []models.ItResourceData) {

	// 数据库连接
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 ItResourcesData 数组并插入到数据库
	for _, resource := range itData {
		// 准备插入SQL语句（使用参数化查询）
		stmt, err := db.Prepare("INSERT INTO it_resource_data (lrid, show_name, share_url, pwd, price, costscore, size, tag) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		// 执行插入操作
		_, err = stmt.Exec(resource.Lrid, resource.ShowName, resource.ShareURL, resource.Pwd, resource.Price, resource.Costscore, resource.Size, resource.Tag)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("ItResourcesData array saved successfully")
}

func InserBookData(bookData []models.BookResourceData) {

	// 数据库连接
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 BookResourcesData 数组并插入到数据库
	for _, resource := range bookData {
		// 准备插入SQL语句（使用参数化查询）
		stmt, err := db.Prepare("INSERT INTO book_resource_data (lr_id, show_name, share_url, pwd) VALUES (?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		// 执行插入操作
		_, err = stmt.Exec(resource.LRID, resource.ShowName, resource.ShareURL, resource.PWD)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("BookResourcesData array saved successfully")
}
