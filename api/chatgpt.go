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
	message string
	user    string
}

func (c *ChatGpt) GetMessage(message string) string {
	data := data{user: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363", message: message}
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Panicln(err)
	}
	resp, err := http.Post("http://0.00000.work/index", "application/json", bytes.NewReader(marshal))
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	fmt.Println(string(r))
	return string(r)

	//strings.ReplaceAll(, "{{end}} ", "\n")
}
