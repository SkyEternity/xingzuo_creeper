package controllers

import (
	"fmt"
	"log"
	"strings"
	"xingzuoCreeper/models"

	"github.com/PuerkitoBio/goquery"
)

var (
	shierBaseUrl = "https://www.xingzuo360.cn"
)

//爬取一个类型的一页已经完成
// https://www.xingzuo360.cn/baiyangzuo/p2.html 列表累加就行

func WorkShier() {
	// getShiClassAll()
	getShiList()
}
func getShiClassAll() {
	err, resp := RequestFn(shierBaseUrl + "/shierxingzuo")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	d.Find(".xz_nav a").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}

type listParams struct {
	Url   string
	Title string
	Desc  string
	Cover string
}

//获取列表界面
func getShiList() {
	err, resp := RequestFn("https://www.xingzuo360.cn/baiyangzuo/")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var listParamsData listParams
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	d.Find(".public_column_list ul li").Each(func(i int, s *goquery.Selection) {
		listParamsData.Url, _ = s.Find("a").Attr("href")
		listParamsData.Title = s.Find(".info_title").Text()
		listParamsData.Desc = s.Find(".info_txt").Text()
		listParamsData.Cover, _ = s.Find("img").Attr("src")
		// if i == 0 {
		getShiDetails(listParamsData)
		// }

	})
}

func getShiDetails(listParamsData listParams) {
	var test models.Shixingzuo
	test.Desc = listParamsData.Desc
	test.Cover = listParamsData.Cover[:strings.Index(listParamsData.Cover, "?")]
	test.Title = listParamsData.Title

	//请求详情页
	err, resp := RequestFn(listParamsData.Url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	test.Content, _ = d.Find(".dc_article_content").Html()
	info := d.Find(".dc_source").Text()
	author := strings.Split(info, "：")[2]
	contentTime := strings.Split(strings.Split(info, "：")[1], " ")[0]
	test.Author = author
	test.ContentTime = contentTime
	test.Query = "十二星座"
	dbShiSave(test)
}

func dbShiSave(data models.Shixingzuo) {
	data.AddData()
}
