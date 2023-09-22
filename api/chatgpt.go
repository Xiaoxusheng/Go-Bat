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
	Message string
}

type data struct {
	Model    string `json:"model,omitempty"`
	Messages []struct {
		Role    string `json:"role,omitempty"`
		Content string `json:"content,omitempty"`
	} `json:"messages,omitempty"`
}

//此功能暂未实现

func (c *ChatGpt) GetMessage(message string) string {

	d := data{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role,omitempty"`
			Content string `json:"content,omitempty"`
		}([]struct {
			Role    string
			Content string
		}{{
			Role:    "user",
			Content: message,
		}}),
	}

	m, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	res, err := http.NewRequest(http.MethodPost, "https://api.chatanywhere.com.cn/v1/chat/completions", strings.NewReader(string(m)))
	if err != nil {
		log.Println(err)
		return ""
	}
	res.Header.Set("Content-Type", "application/json")
	res.Header.Set("Authorization", "Bearer "+config.K.Mode.Key)
	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	cj := new(config.Cj)
	err = json.Unmarshal(b, cj)
	if err != nil {
		return ""
	}
	fmt.Println(string(b), cj.Choices[0].Message.Content)
	return strings.ReplaceAll(cj.Choices[0].Message.Content, "\\", "")

}

// //限制
func (c *ChatGpt) Limit(s config.Messages) string {
	ctx := context.Background()
	id := s.UserId
	if s.UserId == 0 {
		id = s.GroupId
	}
	n, err := config.Rdb.HGet(ctx, "u", strconv.FormatInt(id, 10)).Int()
	fmt.Println("n", n)

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
