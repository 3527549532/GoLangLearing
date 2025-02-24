package controllers

import (
	"_course_pachong/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建一个默认的 Gin 引擎
var r = gin.Default()

func ReceiveItPostRequest() {

	//部署静态html资源
	r.Static("/static", "./static")

	// 定义一个路由和处理函数
	r.POST("/learn/api/itcourselist/", func(c *gin.Context) {
		// 设置返回的内容类型
		c.Header("Content-Type", "application/json")

		// 返回原始的 JSON 字节切片作为响应体，并设置状态码为 200 OK
		c.Data(http.StatusOK, "application/json", dao.SearchIt())
	})

	// 定义一个路由和处理函数
	r.POST("/learn/api/freebooks", func(c *gin.Context) {
		// 设置返回的内容类型
		c.Header("Content-Type", "application/json")

		// 返回原始的 JSON 字节切片作为响应体，并设置状态码为 200 OK
		c.Data(http.StatusOK, "application/json", dao.SearchBook())
	})

	r.Run(":8080")
}

func ReceiveBookPostRequest() {

	// 定义一个路由和处理函数
	r.POST("/learn/api/freebooks", func(c *gin.Context) {
		// 设置返回的内容类型
		c.Header("Content-Type", "application/json")

		// 返回原始的 JSON 字节切片作为响应体，并设置状态码为 200 OK
		c.Data(http.StatusOK, "application/json", dao.SearchBook())
	})
	r.Run(":8080")
}
