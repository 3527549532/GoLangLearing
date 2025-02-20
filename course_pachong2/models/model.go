package models

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

// 定义单个资源列配置项的结构体
type ResourceColumnConfig struct {
	ID    uint    `gorm:"primaryKey"`     // 主键
	Title string  `json:"title"`          // 标题
	Slot  *string `json:"slot,omitempty"` // 可选的slot字段
	Key   *string `json:"key,omitempty"`  // 可选的key字段
}

// 用于JSON输出的结构体
type ResourceColumnConfigOutput struct {
	Title string  `json:"title"`
	Slot  *string `json:"slot,omitempty"`
	Key   *string `json:"key,omitempty"`
}

// 前沿IT学习精品结构体
type ResponseData struct {
	ItResourceColNums interface{} `json:"itresourcecolnums"`
	ItResourceData    interface{} `json:"itresourcedata"`
	ItResourceCount   int         `json:"itresourcecount"`
}

type BookResourceDataResponseData struct {
	BookResourceColNums interface{} `json:"bookresourcecolnums"`
	BookResourceData    interface{} `json:"bookresourcedata"`
	BookResourceCount   int         `json:"bookresourcecount"`
}

// 精品电子书结构体
type BookResourceData struct {
	LRID     int    `json:"lrid"`
	ShowName string `json:"show_name"`
	ShareURL string `json:"share_url"`
	PWD      string `json:"pwd"`
}

// 定义结构体BookResourceDataColumn用于记录精品电子书目录结构
type BookResourceDataColumn struct {
	Title string `json:"title"`
	Slot  string `json:"slot"`
}
