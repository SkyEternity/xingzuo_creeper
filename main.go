package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"xingzuoCreeper/models"

	"github.com/PuerkitoBio/goquery"
)

var (
	baseUrl        = "https://www.d1xz.net"
	datas          []models.Xingzuo
	classInit      []string
	classCountList []int
	wg             sync.WaitGroup
	startTime      time.Time
	endTime        time.Time
)

func requestFn(url string) (err error, resp *http.Response) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	req.Header.Add("Cookie", "Hm_lvt_d96f3af1d84a28ce7598db6f236958bc=1621837279,1622097792,1622099454,1622441576; Hm_lpvt_d96f3af1d84a28ce7598db6f236958bc=1622444901")
	req.Header.Add("Referer", "https://www.d1xz.net")

	resp, err = client.Do(req)
	req.Close = true
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	startTime = time.Now()
	work()
}
func work() {
	fmt.Println("正在爬取.....")
	getClassAll()
	getAllPage()
	endTime = time.Now()
	fmt.Println(endTime.Sub(startTime)) //总爬取时间
}

//获取到全部的分类链接页
func getClassAll() {
	err, resp := requestFn(baseUrl)
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
	err, resp := requestFn(u)
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
	err, resp := requestFn(url)
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
	var data models.Xingzuo
	for i, v := range childUrl {
		wg.Add(1)
		limiter <- true
		go getDetails(baseUrl+v, data, descs[i]) //获取详情页采取并发操作
	}
	wg.Wait()
}
func getDetails(url string, data models.Xingzuo, desc string) {
	data.Desc = desc
	err, resp := requestFn(url)
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

func dbSave(data models.Xingzuo) {
	data.AddData()
}
