package message

import (
	"Go-Bat/Ginterface"
	"Go-Bat/api"
	"Go-Bat/config"
	"context"
	"log"
	"strconv"
	"time"
)

type ZBat struct {
	bat Ginterface.Bat
}

func NewZBat() *ZBat {
	return &ZBat{}
}

// 向外暴露的接口
func (zbat *ZBat) Tactics() {
	if config.K.Mode.Mode == "T" {
		zbat.setStrategy(new(api.Text))
	} else {
		zbat.setStrategy(new(api.Picture))
	}
}

// 设置策略
func (zbat *ZBat) setStrategy(B Ginterface.Bat) {
	zbat.bat = B
}

// 处理函数
func (zbat *ZBat) Deal(mess config.Messages) {
	//redis记录人数
	ctx := context.Background()
	_, err := config.Rdb.HSet(ctx, "chat", strconv.FormatInt(mess.UserId, 10), mess.Message).Result()
	if err != nil {
		log.Println(err)
		return
	}
	config.Rdb.Expire(ctx, "chat", time.Minute*10)
	log.Println("聊天人数：", len(config.Rdb.HGetAll(ctx, "chat").Val()))
	//处理消息
	log.Println("处理消息", mess.Message)
	if zbat.bat == nil {
		return
	}
	//设置
	if mess.Message == "文字模式" {
		zbat.setStrategy(new(api.Text))
		return
	} else if mess.Message == "图片模式" {
		zbat.setStrategy(new(api.Picture))
		return
	}
	zbat.bat.Controls(mess)
}
