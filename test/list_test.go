package test

import (
	"Go-Bat/config"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"testing"
)

func TestList(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:" + strconv.Itoa(config.K.Server.Port) + "/get_friend_list")
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
	list := new(config.Flist)
	err = json.Unmarshal(resp, list)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(list.Data)

}
