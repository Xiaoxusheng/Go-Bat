package api

import (
	"Go-Bat/config"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type Class struct {
	Cookie []*http.Cookie
	w      int64
}

func (c *Class) getCookie() {
	res, err := http.Get("https://passport2.chaoxing.com/api/login?" + "name=" + config.K.ChaoXing.Name + "&pwd=" + config.K.ChaoXing.Password)
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	c.Cookie = res.Cookies()

	//fmt.Println(c.Cookie)
}

func (c *Class) GetClass() string {
	//c.w不存在时，默认为当前周
	if c.w == 0 {
		c.w = int64(math.Ceil(float64((time.Now().Unix() - time.Date(time.Now().Year(), 2, 6, 0, 0, 0, 0, time.Local).Unix()) / (1000 * 60 * 60 * 24))))
		fmt.Println(c.w)
	}
	if c.w < 1 || c.w > 18 {
		return "没课啊,靓仔"
	}
	c.getCookie()
	h := http.Client{}
	req, err := http.NewRequest("GET", "https://kb.chaoxing.com/pc/curriculum/getMyLessons"+"?week="+strconv.FormatInt(c.w, 10), nil)
	if err != nil {
		log.Panicln(err)
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
		log.Panicln(err)
	}
	fmt.Println(string(resp))
	return string(resp)
}
