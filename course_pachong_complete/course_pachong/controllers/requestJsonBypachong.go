package controllers

import (
	"_course_pachong/dao"
	"_course_pachong/models"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
	"strings"
)

func RequestItJsonDataBypachong() {
	url := "https://dl.zzyyww.cn/learn/api/itcourselist/"
	method := "POST"

	payload := strings.NewReader(`{"learn_type":"type_it","key1_1":"","key1_2":"","key1_3":"","it_current":15,"sel_ittags":[]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Origin", "https://dl.zzyyww.cn")
	req.Header.Add("Referer", "https://dl.zzyyww.cn/learn/")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not(A:Brand\";v=\"99\", \"Google Chrome\";v=\"133\", \"Chromium\";v=\"133\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析JSON数据
	var temp struct {
		Itresourcedata []models.ItResourceData `json:"itresourcedata"`
	}

	err = json.Unmarshal(body, &temp)
	if err != nil {
		log.Fatalf("解析JSON数据失败: %v", err)
	}

	dao.InserItData(temp.Itresourcedata)

}

func RequestBookJsonDataBypachong() {
	url := "https://dl.zzyyww.cn/learn/api/freebooks"
	method := "POST"

	payload := strings.NewReader(`{"key3_1":"","key3_2":"","key3_3":"","book_current":15,"sel_booktags":[]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Origin", "https://dl.zzyyww.cn")
	req.Header.Add("Referer", "https://dl.zzyyww.cn/learn/")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not(A:Brand\";v=\"99\", \"Google Chrome\";v=\"133\", \"Chromium\";v=\"133\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析JSON数据
	var temp struct {
		Bookresourcedata []models.BookResourceData `json:"bookresourcedata"`
	}

	err = json.Unmarshal(body, &temp)
	if err != nil {
		log.Fatalf("解析JSON数据失败: %v", err)
	}

	dao.InserBookData(temp.Bookresourcedata)

}

func RequestKlJsonDataBypachong() {
	url := "https://dl.zzyyww.cn/learn/api/knowledge/"
	method := "POST"

	payload := strings.NewReader(`{"learn_type":"type_knowledge","key2_1":"","key2_2":"","key2_3":"","kl_current":15,"sel_kltags":[]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Origin", "https://dl.zzyyww.cn")
	req.Header.Add("Referer", "https://dl.zzyyww.cn/learn/")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not(A:Brand\";v=\"99\", \"Google Chrome\";v=\"133\", \"Chromium\";v=\"133\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析JSON数据
	var temp struct {
		Klresourcedata []models.KnowledgeResourcesData `json:"klresourcedata"`
	}

	err = json.Unmarshal(body, &temp)
	if err != nil {
		log.Fatalf("解析JSON数据失败: %v", err)
	}

	dao.InserKlData(temp.Klresourcedata)

}
