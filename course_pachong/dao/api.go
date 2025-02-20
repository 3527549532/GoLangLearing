package dao

import (
	"_course_pachong/models"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func SearchIt() []byte {
	// 连接到MySQL数据库
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 查询所有ResourceColumnConfig记录
	var configs []models.ResourceColumnConfig
	result := db.Find(&configs)
	if result.Error != nil {
		log.Fatalf("failed to query database: %v", result.Error)
	}

	// 过滤并转换数据
	var outputConfigs []models.ResourceColumnConfigOutput
	for _, config := range configs {
		if config.Slot != nil || config.Key != nil { // 只包含至少有一个非nil字段的记录
			outputConfig := models.ResourceColumnConfigOutput{
				Title: config.Title,
				Slot:  config.Slot,
				Key:   config.Key,
			}
			outputConfigs = append(outputConfigs, outputConfig)
		}
	}

	//查询所有ItResourceData记录
	var itresourcedata []models.ItResourceData
	itresourcedataResult := db.Find(&itresourcedata)
	totalCount := len(itresourcedata)
	if itresourcedataResult.Error != nil {
		log.Fatalf("failed to query database: %v", itresourcedataResult.Error)
	}
	response := models.ResponseData{
		ItResourceColNums: outputConfigs,
		ItResourceData:    itresourcedata,
		ItResourceCount:   totalCount,
	}
	// 序列化为JSON格式
	jsonData, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("failed to marshal JSON: %v", err)
	}

	return jsonData
}

func SearchBook() []byte {
	// 连接到MySQL数据库
	dsn := "root:1234@tcp(127.0.0.1:3306)/pachong_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	//查询所有BookResourceData记录
	var BookResourceData []models.BookResourceData
	BookResourceDataResult := db.Find(&BookResourceData)
	totalCount := len(BookResourceData)
	if BookResourceDataResult.Error != nil {
		log.Fatalf("failed to query database: %v", BookResourceDataResult.Error)
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
