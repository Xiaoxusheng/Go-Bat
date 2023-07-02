package abstraction

import (
	"Go-Bat/api"
	"Go-Bat/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// 抽象接口
type Bat interface {
	Controls(s any)
}

type Data struct {
	User_id      int64  `json:"user_id"`
	Group_id     int64  `json:"group_id"`
	Message      string `json:"message"`
	Message_type string `json:"message_type"`
	Auto_escape  bool   `json:"auto_escape"`
}

type GoBat struct {
	bat Bat
}

// 设置策略
func (bat *GoBat) setStrategy(B Bat) {
	bat.bat = B
}

func (bat *GoBat) Send(d Data) {
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
	ctx := context.Background()
	config.Rdb.Set(ctx, strconv.FormatInt(mess.User_id, 10), mess.Message, time.Minute*10)
	fmt.Println("聊天人数：", len(config.Rdb.Keys(ctx, "*").Val()))
	go func() {
		for {
			select {
			case c := <-config.SendChan:
				fmt.Println("数据进入chan")
				//定时消息
				if mess.User_id == 0 {
					mess.User_id = 3096407768
					mess.Message_type = "private"
				}
				bat.Send(Data{User_id: mess.User_id, Message: c, Message_type: mess.Message_type, Auto_escape: false, Group_id: mess.Group_id})
			default:
				continue
			}
		}
	}()
	//处理消息
	bat.bat.Controls(mess)

}

// 向外暴露的接口
func (bat *GoBat) Tactics() {
	if config.K.Mode.Mode == "T" {
		bat.setStrategy(new(api.Text))
	} else {
		bat.setStrategy(new(api.Picture))
	}
}
