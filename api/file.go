package api

import (
	"Go-Bat/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type File struct {
	UserId int64  `json:"user_id"`
	File   string `json:"file"`
	Name   string `json:"name"`
}

func (f *File) Upload() {
	f.UserId = config.K.Bat.QQ
	f.Name = time.Now().Format(time.DateOnly) + ".log"
	f.File = "/root/GoBatRoot/log/" + time.Now().Format(time.DateOnly) + ".log"
	marshal, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.Post("http://127.0.0.1:"+strconv.Itoa(config.K.Server.Port)+"/upload_private_file", "application/json", bytes.NewReader(marshal))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))

}
