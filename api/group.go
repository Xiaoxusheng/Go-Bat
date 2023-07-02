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
	Group_id int64  `json:"group_id"`
	User_id  int64  `json:"user_id"`
	Duration uint32 `json:"duration"`
	File     string `json:"file,omitempty"`
	Name     string `json:"name,omitempty"`
	Folder   string `json:"folder,omitempty"`
}

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
	g.ban(GroupList{s.(config.Messages).Group_id, parseInt, uint32(time), "", "", ""})

}

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
