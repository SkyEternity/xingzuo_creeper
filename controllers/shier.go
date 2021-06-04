package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
	"xingzuoCreeper/models"

	"github.com/PuerkitoBio/goquery"
)

var (
	shierBaseUrl = "https://www.xingzuo360.cn"
	countArr     []int
	wgshi        sync.WaitGroup
	sTime        time.Time
	eTime        time.Time
)

//爬取一个类型的一页已经完成
// https://www.xingzuo360.cn/baiyangzuo/p2.html 列表累加就行

func WorkShier() {
	sTime = time.Now()
	fmt.Println("正在爬取.....")
	getShiClassAll()
	eTime = time.Now()
	fmt.Println(eTime.Sub(sTime)) //总爬取时间
	// getShiList()
}
func getShiClassAll() {
	err, resp := RequestFn(shierBaseUrl + "/shierxingzuo")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	d.Find(".xz_nav a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		getShiCountRequest(url) //再去获取每类的总页数
	})
}
func getShiCountRequest(url string) {
	err, resp := RequestFn(url + "p1.html")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	count, _ := strconv.Atoi(d.Find(".next").Prev().Text())
	for i := 1; i <= count; i++ {
		getShiList(url + "p" + strconv.Itoa(i) + ".html")
	}
}

//总页数拿到了然后去请求页面
type listParams struct {
	Url   string
	Title string
	Desc  string
	Cover string
}

//控制并发数

var limitShi = make(chan bool, 20)

//获取列表界面
func getShiList(url string) {
	err, resp := RequestFn(url)
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
		wgshi.Add(1)
		limitShi <- true
		go getShiDetails(listParamsData)
	})
	wgshi.Wait()
}

func getShiDetails(listParamsData listParams) {
	var test models.Shixingzuo
	test.Desc = listParamsData.Desc
	if listParamsData.Cover != "" {
		test.Cover = listParamsData.Cover[:strings.Index(listParamsData.Cover, "?")]
	}
	test.Title = listParamsData.Title
	//请求详情页
	err, resp := RequestFn(listParamsData.Url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	test.Content, _ = d.Find(".detail_box .dc_article_content").Html()
	info := d.Find(".dc_source").Text()
	author := strings.Split(info, "：")[2]
	contentTime := strings.Split(strings.Split(info, "：")[1], " ")[0]
	test.Author = author
	test.ContentTime = contentTime
	test.Query = "十二星座"
	dbShiSave(test)
	defer wgshi.Done()
	time.Sleep(100 * time.Millisecond)
	<-limitShi
}

func dbShiSave(data models.Shixingzuo) {
	data.AddData()
}
