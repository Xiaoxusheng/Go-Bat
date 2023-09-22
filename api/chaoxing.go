package api

import (
	"Go-Bat/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type class struct {
	Cookie []*http.Cookie
	w      int64
}

func (c *class) getCookie() {
	res, err := http.Get("https://passport2.chaoxing.com/api/login?" + "name=" + config.K.ChaoXing.Name + "&pwd=" + config.K.ChaoXing.Password)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	c.Cookie = res.Cookies()

	//fmt.Println(c.Cookie)
}

func (c *class) GetClass() string {
	//c.w不存在时，默认为当前周
	if c.w == 0 {
		c.w = int64(math.Ceil(float64((time.Now().Unix()-time.Date(time.Now().Year(), 9, 1, 0, 0, 0, 0, time.Local).Unix())/(7*60*60*24)))) + 1
		fmt.Println(c.w)
	}
	if c.w > 10 {
		return "没课啊,靓仔"
	}
	c.getCookie()
	h := http.Client{}
	req, err := http.NewRequest("GET", "https://kb.chaoxing.com/pc/curriculum/getMyLessons"+"?week="+strconv.FormatInt(c.w, 10), nil)
	if err != nil {
		log.Println(err)
	}
	for _, cookie := range c.Cookie {
		req.AddCookie(cookie)
	}
	res, err := h.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(resp))
	class := config.Class{}
	err = json.Unmarshal(resp, &class)
	if err != nil {
		log.Println("json解析失败")
		return ""
	}
	str := "第" + strconv.Itoa(class.Data.Curriculum.CurrentWeek) + "周课表\n" + "本周共[" + strconv.Itoa(class.Data.Curriculum.CurriculumCount) + "]节课\n"

	for i := 0; i < len(class.Data.LessonArray); i++ {
		str += "星期" + strconv.Itoa(class.Data.LessonArray[i].DayOfWeek) + "\n" + "课程名称: " + class.Data.LessonArray[i].ClassName + "\n" + "上课地点: " + class.Data.LessonArray[i].Location + "\n" +
			"节次：" + strconv.Itoa(class.Data.LessonArray[i].BeginNumber) + "--" + strconv.Itoa(class.Data.LessonArray[i].BeginNumber+1) + "\n"

	}
	return str
}

// 定时通知
func (c *class) set() string {
	//c.w不存在时，默认为当前周
	if c.w == 0 {
		c.w = int64(math.Ceil(float64((time.Now().Unix()-time.Date(time.Now().Year(), 9, 4, 0, 0, 0, 0, time.Local).Unix())/(7*60*60*24)))) + 1
	}
	c.getCookie()
	h := http.Client{}
	req, err := http.NewRequest("GET", "https://kb.chaoxing.com/pc/curriculum/getMyLessons"+"?week="+strconv.FormatInt(c.w, 10), nil)
	if err != nil {
		log.Println(err)
	}
	for _, cookie := range c.Cookie {
		req.AddCookie(cookie)
	}
	res, err := h.Do(req)
	if err != nil {
		log.Println("请求出错")
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(resp))
	class := config.Class{}
	err = json.Unmarshal(resp, &class)
	if err != nil {
		log.Println("json解析失败")
	}
	str := ""
	for i := 0; i < len(class.Data.LessonArray); i++ {
		if int(time.Now().Weekday()) == class.Data.LessonArray[i].DayOfWeek {
			str += "星期" + strconv.Itoa(class.Data.LessonArray[i].DayOfWeek) + "\n" + "课程名称: " + class.Data.LessonArray[i].ClassName + "\n" + "上课地点: " + class.Data.LessonArray[i].Location + "\n" +
				"节次：" + strconv.Itoa(class.Data.LessonArray[i].BeginNumber) + "--" + strconv.Itoa(class.Data.LessonArray[i].BeginNumber+1) + "\n"
		}
	}
	return str
}

func (c *class) SetTime() {
	log.Println("定时任务启动中")
	t1 := time.Now()
	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute()+1, 0, 0, t1.Location())
	log.Println("任务在" + t2.Sub(t1).String() + "秒后启动")
	t3 := time.NewTicker(t2.Sub(t1))
	for {
		select {
		case <-t3.C:
			config.SendChan <- config.SendMessage{
				UserId:      3096407768,
				Message:     c.set(),
				MessageType: "",
				AutoEscape:  false,
			}
			t1 = time.Now()
			t2 = time.Date(t1.Year(), t1.Month(), t1.Day()+1, 7, 0, 0, 0, t1.Location())
			t3 = time.NewTicker(t2.Sub(t1))
		}
	}

}
