package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ChatGpt struct {
	Message string
}

type data struct {
	text            string
	parentMessageId string
}

func (c *ChatGpt) GetMessage(message string) string {
	h := &http.Client{}
	data := data{parentMessageId: "chatcmpl-7UQSv9qcxBvZsJoByHGFWeEyJI9zV", text: message}
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Panicln(err)
	}
	res, err := http.NewRequest("POST", "https://openai.ppjun.cn/chat", bytes.NewReader(marshal))
	res.Header.Set("Origin", "https://openai.ppjun.cn")
	res.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.51")
	res.Header.Set("Referer", "https://openai.ppjun.cn/")
	//res, err := http.Post("https://openai.ppjun.cn/chat", "application/json", bytes.NewReader(marshal))
	if err != nil {
		log.Panicln(err)
	}
	resp, err := h.Do(res)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	fmt.Println(string(r))
	return string(r)

	//strings.ReplaceAll(, "{{end}} ", "\n")
}
