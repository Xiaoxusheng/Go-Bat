package test

import (
	"Go-Bat/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
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

func TestClass(t *testing.T) {
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
			Content: "Hello!",
		}}),
	}
	fmt.Println(d)
	m, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
	}
	res, err := http.NewRequest(http.MethodPost, "https://api.chatanywhere.com.cn/v1/chat/completions", bytes.NewReader(m))
	if err != nil {
		log.Println(err)
	}
	res.Header.Set("Content-Type", "application/json")
	res.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	res.Header.Set("Authorization", "Bearer sk-H2Ea8g7of8MmOeu402a14ULPWuhijzH9zkGEq3KBXDdhEfeb")
	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	cj := new(config.Cj)
	err = json.Unmarshal(b, cj)
	if err != nil {
	}
	fmt.Println(string(b), "\n", cj.Choices[0].Message.Content)
	fmt.Println(resp.StatusCode)

}
