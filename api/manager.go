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

type manager struct {
	Data struct {
		MessageType string `json:"message_type,omitempty"`
		Message     string `json:"message,omitempty"`
	}
	Status string `json:",omitempty"`
}

/*res {"data":{"group":false,"message":"r","message_id":1104839189,"message_id_v2":"00000000b88f6ed800002a90",
"message_seq":10896,"message_type":"private","real_id":10896,"sender":{"nickname":"Ra","user_id":3096407768},
"time":1687999849},"message":"","retcode":0,"status":"ok"}*/

var M = manager{}

// 防撤回
func (m *manager) preventRecall(c config.Messages) {
	res, err := http.Get("http://127.0.0.1:5000/get_msg" + "?message_id=" + strconv.FormatInt(c.MessageId, 10))
	if err != nil {
		log.Panicln(err)
	}
	resp, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("res", string(resp))

	err = json.Unmarshal(resp, &M)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("m", M)

}
