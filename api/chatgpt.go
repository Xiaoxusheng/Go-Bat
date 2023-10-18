package api

import (
	"Go-Bat/config"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ChatGpt struct {
	Model    string `json:"model,omitempty"`
	Messages []struct {
		Role    string `json:"role,omitempty"`
		Content string `json:"content,omitempty"`
	} `json:"messages,omitempty"`
	Temperature float64 `json:"temperature"`
	time        time.Time
}

// 定义全局变量
var d = &ChatGpt{Model: "gpt-3.5-turbo",
	Messages: []struct {
		Role    string `json:"role,omitempty"`
		Content string `json:"content,omitempty"`
	}([]struct {
		Role    string
		Content string
	}{{
		Role:    "user",
		Content: "",
	}}),
	Temperature: 0.8}

var chat = make(map[int64]*ChatGpt, 6)

func (c *ChatGpt) GetMessage(id int64, message string) string {
	//如果id没有对应的切片
	if _, ok := chat[id]; !ok {
		d.time = time.Now()
		d.Messages[0].Content = message
		chat[id] = d
	}
	//添加
	chat[id].Messages = append(chat[id].Messages, struct {
		Role    string `json:"role,omitempty"`
		Content string `json:"content,omitempty"`
	}(struct {
		Role    string
		Content string
	}{
		Role:    "user",
		Content: message,
	}))

	m, err := json.Marshal(d)
	if err != nil {
		fmt.Println("json", err)
		return err.Error()
	}
	res, err := http.NewRequest(http.MethodPost, "https://api.chatanywhere.com.cn/v1/chat/completions", strings.NewReader(string(m)))
	if err != nil {
		return err.Error()
	}
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Authorization", "Bearer "+config.K.Mode.Key)
	client := &http.Client{
		//Timeout: time.Second * 10,
	}
	resp, err := client.Do(res)
	if err != nil {
		log.Println(err)
		return strings.ReplaceAll(err.Error(), "https://api.chatanywhere.com.cn/v1/chat/completions", " ")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	if resp.StatusCode == 429 {
		return errors.New("使用超过限制").Error()
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	fmt.Println(string(b), resp.StatusCode)
	cj := new(config.Cj)
	err = json.Unmarshal(b, cj)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	//最多记录6条上下文记录，超过就删除  停留30 删除时间为30分钟
	fmt.Println(chat[id].Messages, len(chat[id].Messages), time.Now().Sub(chat[id].time) > 30)

	if len(chat[id].Messages) > 6 || time.Now().Sub(chat[id].time).Minutes() > 30 {
		//大于120分钟
		if time.Now().Sub(chat[id].time).Minutes() > 120 {
			//直接删除
			delete(chat, id)
		}
		chat[id].Messages = []struct {
			Role    string `json:"role,omitempty"`
			Content string `json:"content,omitempty"`
		}{}
	}

	chat[id].Messages = append(chat[id].Messages, struct {
		Role    string `json:"role,omitempty"`
		Content string `json:"content,omitempty"`
	}(struct {
		Role    string
		Content string
	}{
		Role:    cj.Choices[0].Message.Role,
		Content: cj.Choices[0].Message.Content,
	}))
	//更新时间
	chat[id].time = time.Now()
	fmt.Println(d.Messages)
	return strings.ReplaceAll(cj.Choices[0].Message.Content, "\\", "")
}

// 限制
func (c *ChatGpt) Limit(s config.Messages) string {
	ctx := context.Background()
	id := s.UserId
	if s.UserId == 0 {
		id = s.GroupId
	}
	n, err := config.Rdb.HGet(ctx, "u", strconv.FormatInt(id, 10)).Int()
	fmt.Println("n", n)
	fmt.Println(s.UserId != config.K.Bat.QQ, s.UserId, config.K.Bat.QQ)
	if s.UserId != config.K.Bat.QQ && n > 20 {
		log.Println("1", err)
		return errors.New("今日次数用完！").Error()
	}
	//判断hash是否存在，不存在再进行添加
	if !config.Rdb.HExists(ctx, "u", strconv.FormatInt(id, 10)).Val() {
		_, err = config.Rdb.HSet(ctx, "u", strconv.FormatInt(id, 10), 1).Result()
		if err != nil {
			log.Println("2", err)
			return errors.New("出错了").Error()
		}
		config.Rdb.Expire(ctx, "u", time.Hour*24)
	}
	//累计加1
	_, err = config.Rdb.HIncrBy(ctx, "u", strconv.FormatInt(id, 10), 1).Result()
	if err != nil {
		log.Println("3", err)
		return errors.New("出错了").Error()
	}
	return ""
}
