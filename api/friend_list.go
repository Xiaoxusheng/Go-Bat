package api

import (
	"Go-Bat/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type FriendList struct {
	Data []struct {
		Nickname string `json:"nickname"`
		Remark   string `json:"remark"`
		UserId   int64  `json:"user_id"`
	} `json:"data"`
	Message string `json:"message"`
	Retcode int    `json:"retcode"`
	Status  string `json:"status"`
}

func (f *FriendList) GetFriendList() {
	res, err := http.Get("http://127.0.0.1:" + strconv.Itoa(config.K.Server.Port) + "/get_unidirectional_friend_list")
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	list := new(FriendList)
	err = json.Unmarshal(resp, list)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(list.Data)
}
