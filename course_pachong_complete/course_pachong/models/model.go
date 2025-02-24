package models

// It 数据结构体
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

// 获取 It 页的列标题
type ItResourceDataColumn struct {
	ID    uint    `gorm:"primaryKey"`     // 主键
	Title string  `json:"title"`          // 标题
	Slot  *string `json:"slot,omitempty"` // 可选的slot字段
	Key   *string `json:"key,omitempty"`  // 可选的key字段
}

// 前沿 IT 学习精品转 Json 结构体
type ResponseItData struct {
	ItResourceColNums interface{}      `json:"itresourcecolnums"`
	ItResourceData    []ItResourceData `json:"itresourcedata"`
	ItResourceCount   int64            `json:"itresourcecount"`
}

// 精品电子书结构体
type BookResourceData struct {
	LRID     int    `json:"lrid"`
	ShowName string `json:"show_name"`
	ShareURL string `json:"share_url"`
	PWD      string `json:"pwd"`
}

// 精品电子书列标签结构体
type BookResourceDataColumn struct {
	Title string `json:"title"`
	Slot  string `json:"slot"`
}

// 精品电子书转 Json 结构体
type BookResourceDataResponseData struct {
	BookResourceColNums interface{} `json:"bookresourcecolnums"`
	BookResourceData    interface{} `json:"bookresourcedata"`
	BookResourceCount   int64       `json:"bookresourcecount"`
}

// 精品知识付费结构体
type KnowledgeResourcesData struct {
	Lrid      int    `json:"lrid"`
	ShowName  string `json:"show_name"`
	ShareURL  string `json:"share_url"`
	Pwd       string `json:"pwd"`
	Price     string `json:"price"` // 根据数据，这里的价格是字符串格式，如果需要可以改为float64并处理转换
	CostScore int    `json:"costscore"`
	Size      string `json:"size"` // 根据数据，这里的大小是字符串格式，如果需要可以改为适当的数据类型并处理转换
	Tag       string `json:"tag"`
}

// 精品知识付费列标签结构体
type KnowledgeResourceDataColumn struct {
	Title string `json:"title"`
	Slot  string `json:"slot"`
}

// 精品知识付费转 Json 结构体
type KnowledgeResourceDataResponseData struct {
	KnowledgeResourceColNums interface{} `json:"klresourcecolnums"`
	KnowledgeResourceData    interface{} `json:"klresourcedata"`
	KnowledgeResourceCount   int64       `json:"klresourcecount"`
}

// 定义用于获取IT请求体json数据的结构体
type RequestITBodyData struct {
	Key3_1    string   `json:"key3_1"`
	Key3_2    string   `json:"key3_2"`
	Key3_3    string   `json:"key3_3"`
	ITCurrent int      `json:"it_current"`
	SelITtags []string `json:"sel_ittags"`
}

// 定义用于获取 book 请求体json数据的结构体
type RequestBookBodyData struct {
	Key3_1      string   `json:"key3_1"`
	Key3_2      string   `json:"key3_2"`
	Key3_3      string   `json:"key3_3"`
	BookCurrent int      `json:"book_current"`
	SelBooktags []string `json:"sel_booktags"`
}

// 定义用于获取 Knowledge 请求体json数据的结构体
type RequestKnowledgeData struct {
	Key3_1    string   `json:"key3_1"`
	Key3_2    string   `json:"key3_2"`
	Key3_3    string   `json:"key3_3"`
	KlCurrent int      `json:"kl_current"`
	SelKltags []string `json:"sel_kltags"`
}
