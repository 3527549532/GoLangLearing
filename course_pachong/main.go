package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql" // 注意这里的下划线，表示只导入包但不使用它
)

// 定义结构体来映射JSON数据
type ItResourceData struct {
	Lrid      int    `json:"lrid"`
	ShowName  string `json:"show_name"`
	ShareURL  string `json:"share_url"`
	Pwd       string `json:"pwd"`
	Price     string `json:"price"`
	Costscore int    `json:"costscore"`
	Size      string `json:"size"`
	Tag       string `json:"tag"`
}

type ItCourseList struct {
	Itresourcedata  []ItResourceData `json:"itresourcedata"`
	Itresourcecount int              `json:"itresourcecount"`
}

func main() {
	//静态资源启动

	// 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 设置静态文件目录
	r.Static("/static", "./static")

	// 加载 HTML 模板
	r.LoadHTMLGlob("templates/*")

	// 定义一个路由来服务 HTML 页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 启动服务器
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

	//静态资源配置结束

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
