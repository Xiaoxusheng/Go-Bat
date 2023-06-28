package abstraction

import (
	"Go-Bat/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// 抽象接口
type Bat interface {
	Controls(s any)
}

type Data struct {
	User_id      int64  `json:"user_id"`
	Message      string `json:"message"`
	Message_type string `json:"message_type"`
	Auto_escape  bool   `json:"auto_escape"`
}

type GoBat struct {
	bat Bat
}

// 设置策略
func (bat *GoBat) SetStrategy(B Bat) {
	bat.bat = B
}

func (b *GoBat) Send(d Data) {
	fmt.Println("发送", d)
	marshal, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://127.0.0.1:"+strconv.Itoa(config.K.Server.Port)+"/send_msg", "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		panic(err)
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("res", string(all))
	defer resp.Body.Close()
}

// 调用
func (bat *GoBat) Deal(mess config.Messages) {
	//redis记录人数
	bat.bat.Controls(mess)

	go func() {
		for {
			select {
			case c := <-config.SendChan:
				fmt.Println("数据进入")
				bat.Send(Data{User_id: mess.User_id, Message: c, Message_type: mess.Message_type, Auto_escape: false})
			default:
				continue
			}
		}
	}()
}
