package controllers

import (
	"_course_pachong/dao"
	"_course_pachong/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建一个默认的 Gin 引擎
var R = gin.Default()

func ReceiveItPostRequest() {

	// 定义一个路由和处理函数
	R.POST("/learn/api/itcourselist/", func(c *gin.Context) {

		var jsonData models.RequestITBodyData
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 返回原始的 JSON 字节切片作为响应体，并设置状态码为 200 OK
		c.Data(http.StatusOK, "application/json", dao.SearchIt(jsonData))
	})
}

func ReceiveBookPostRequest() {

	// 定义一个路由和处理函数
	R.POST("/learn/api/freebooks/", func(c *gin.Context) {
		// 设置返回的内容类型
		c.Header("Content-Type", "application/json")

		var jsonData models.RequestBookBodyData
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 返回原始的 JSON 字节切片作为响应体，并设置状态码为 200 OK
		c.Data(http.StatusOK, "application/json", dao.SearchBook(jsonData))
	})
}

func ReceiveKnowledgePostRequest() {

	// 定义一个路由和处理函数
	R.POST("/learn/api/knowledge/", func(c *gin.Context) {
		// 设置返回的内容类型
		c.Header("Content-Type", "application/json")

		var jsonData models.RequestKnowledgeData
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 返回原始的 JSON 字节切片作为响应体，并设置状态码为 200 OK
		c.Data(http.StatusOK, "application/json", dao.SearchKnowledge(jsonData))
	})
}
