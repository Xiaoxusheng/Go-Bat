package api

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"log"
	"os"
	"strings"
	"time"
)

type collyBaidu struct {
}

// 获取百度热搜
func (cl *collyBaidu) crawler() string {
	var str string
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.54"),
		colly.Debugger(&debug.LogDebugger{}),
	)

	c.OnError(func(r *colly.Response, err error) {
		log.Println("任务出现错误-->:", err)
	})
	//爬取内容
	c.OnHTML("div[class='category-wrap_iQLoo horizontal_1eKyQ']", func(e *colly.HTMLElement) {

		str += "标题：" + e.DOM.Find(".c-single-text-ellipsis").Text() + "\t\t" + "内容：" + strings.Replace(e.DOM.Find("div[class='hot-desc_1m_jR large_nSuFU ']").Text(), " 查看更多>", "", 1) + "\t" + "热度：" + e.DOM.Find(".hot-index_1Bl1a").Text() + "\n"

	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	err := c.Visit("https://top.baidu.com/board?tab=realtime")
	if err != nil {
		log.Println(err)
		return ""
	}
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("请求之前回调:", r.URL.String())
	})
	file, err := os.OpenFile("./config/"+time.Now().Format("2006-01-02")+".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Panicln("创建hot.txt出错" + err.Error())
	}
	defer file.Close()
	_, err = file.Write([]byte(str))
	if err != nil {
		log.Panicln("写入文件出错" + err.Error())
	}
	return str

}
