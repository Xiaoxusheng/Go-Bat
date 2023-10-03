package test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func getCookie() []*http.Cookie {
	res, err := http.PostForm("https://service.wut.edu.cn/login", url.Values{
		"uname": {"20204221338"},
		"pwd":   {"bGVpMTI1NjA4"},
		"type":  {"1"},
	})
	if err != nil {
		log.Panicln(err)
	}
	defer res.Body.Close()
	fmt.Println(res.Cookies())
	return res.Cookies()

	//fmt.Println(c.Cookie)
}

func TestKe(t *testing.T) {
	formData := url.Values{}
	formData.Set("type", "1")
	formData.Set("zc", "4")
	client := http.Client{}

	r, err := http.NewRequest(http.MethodPost, "http://jwxt.wut.edu.cn/admin/getXsdSykb", strings.NewReader(formData.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Println(err)
		return
	}
	for _, cookie := range getCookie() {
		r.AddCookie(cookie)
	}
	fmt.Println("发起请求")
	res, err := client.Do(r)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(resp), res.StatusCode)

}
