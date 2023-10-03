package api

import (
	"Go-Bat/config"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type AutoFriend struct {
	Flag    string `json:"flag,omitempty"`
	Approve bool   `json:"approve,omitempty"`
	Remark  string `json:"remark,omitempty"`
}

// 自动舔加好友
// {"post_type":"request","request_type":"friend","time":1679386108,"self_id":2673893724,"user_id":1978150028,"comment":"信息","flag":"1679386108000000"}
func (a *AutoFriend) auto(f config.Messages) {
	marshal, err := json.Marshal(AutoFriend{Flag: f.Flag, Approve: true, Remark: f.Remark})
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.Post("http://127.0.0.1:"+strconv.Itoa(config.K.Server.Port)+"/upload_group_file", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
}
