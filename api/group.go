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
	"strings"
	"unicode"
)

type GroupList struct {
	GroupId  int64  `json:"group_id"`
	UserId   int64  `json:"user_id"`
	Duration uint32 `json:"duration"`
	File     string `json:"file,omitempty"`
	Name     string `json:"name,omitempty"`
	Folder   string `json:"folder,omitempty"`
}

type GroupNotice struct {
	GroupId int64  `json:"group_id"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

// 禁言
func (g *GroupList) receive(s any) {
	//默认为1小时
	time := 60 * 60
	id := s.(config.Messages).Message[strings.Index(s.(config.Messages).Message, "=")+1 : strings.Index(s.(config.Messages).Message, "]")]
	parseInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Panicln(err)
	}
	i := ""
	r := strings.Split(s.(config.Messages).Message, "]")[1]
	for _, v := range r {
		if unicode.IsNumber(v) {
			i += string(v)
		}
	}

	if i != "" {
		parseUint, err := strconv.Atoi(i)
		if err != nil {
			log.Panicln(err)
		}
		time = 60 * parseUint
	}
	fmt.Println("时间", time)
	g.ban(GroupList{s.(config.Messages).GroupId, parseInt, uint32(time), "", "", ""})

}

// 解除
func (g *GroupList) ban(Group GroupList) {
	marshal, err := json.Marshal(Group)
	if err != nil {
		log.Panicln(err)
	}
	resp, err := http.Post("http://127.0.0.1:"+strconv.Itoa(config.K.Server.Port)+"/set_group_ban", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(res))
}

// 群公告
func (g *GroupList) GroupNotice(messages config.Messages) {
	//截取
	n := GroupNotice{
		GroupId: messages.GroupId,
		Content: "",
		Image:   "",
	}
	Notice, err := json.Marshal(n)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.Post("http://127.0.0.1:/_send_group_notice", "application/json", bytes.NewBuffer(Notice))
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(res))
}
