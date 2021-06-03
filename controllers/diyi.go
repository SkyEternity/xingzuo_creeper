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
	baseUrl        = "https://www.d1xz.net"
	datas          []models.Dyxingzuo
	classInit      []string
	classCountList []int
	wg             sync.WaitGroup
	startTime      time.Time
	endTime        time.Time
)

func WorkDy() {
	startTime = time.Now()
	fmt.Println("正在爬取.....")
	getClassAll()
	getAllPage()
	endTime = time.Now()
	fmt.Println(endTime.Sub(startTime)) //总爬取时间
}

//获取到全部的分类链接页
func getClassAll() {
	err, resp := RequestFn(baseUrl)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	d.Find(".constellation li").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		firstPageUrl := baseUrl + href + "list_1.aspx" //十二星座所有第一页的url
		getCountRequest(firstPageUrl)                  //再去获取每类的总页数
	})
}
func getCountRequest(u string) {
	err, resp := RequestFn(u)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	count, _ := strconv.Atoi(d.Find(".page .next").Prev().Text())
	classCountList = append(classCountList, count)
	classInit = append(classInit, u[:strings.Index(u, "_")+1]) //将链接处理成https://www.d1xz.net/astro/Aquarius/list_以便后边的页数进行拼接
}

//定义最大并发数
var limiter = make(chan bool, 20)

//拿到所有类型的所有页面
func getAllPage() {
	for i := 0; i < len(classCountList); i++ {
		for j := 1; j <= classCountList[i]; j++ {
			url := classInit[i] + strconv.Itoa(i) + ".aspx"
			analyseList(url) //请求列表界面
		}
	}
}

//先分析列表页
func analyseList(url string) {
	err, resp := RequestFn(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var childUrl []string
	var descs []string
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	doc.Find(".words_list_ui li").Each(func(i int, s *goquery.Selection) {
		desc := strings.TrimSpace(s.Find(".txt p").Text())
		detailsUrl, _ := s.Find("a").Attr("href")
		descs = append(descs, desc)
		childUrl = append(childUrl, detailsUrl)
	})
	detailsWork(childUrl, descs)
}
func detailsWork(childUrl, descs []string) {
	var data models.Dyxingzuo
	for i, v := range childUrl {
		wg.Add(1)
		limiter <- true
		go getDetails(baseUrl+v, data, descs[i]) //获取详情页采取并发操作
	}
	wg.Wait()
}
func getDetails(url string, data models.Dyxingzuo, desc string) {
	data.Desc = desc
	err, resp := RequestFn(url)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := goquery.NewDocumentFromReader(resp.Body)
	//删除图片中无用的属性
	d.Find(".common_det_con p img").RemoveAttr("title")
	d.Find(".common_det_con p img").RemoveAttr("alt")
	data.Title = d.Find("h1").Text()
	data.Cover, _ = d.Find(".common_det_con p img").Attr("src")
	data.Author = strings.Split(d.Find(".source p span").Eq(1).Text(), "：")[1]
	data.Query = d.Find(".art_con_left .source p span").Eq(2).Text()
	data.Content, _ = d.Find(".common_det_con").Html()
	data.ContentTime = d.Find(".art_con_left .source p span").Eq(0).Text()
	//数据入库
	dbSave(data)
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	<-limiter
}

func dbSave(data models.Dyxingzuo) {
	data.AddData()
}
