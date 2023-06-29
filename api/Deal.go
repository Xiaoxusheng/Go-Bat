package api

import (
	"Go-Bat/config"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type PrivateText struct {
	class   Class
	chatgpt ChatGpt
}

type PrivatePicture struct {
	class   Class
	picture config.Picture
	m       Manager
	t       Timing
}

type GroupText struct {
}

type GroupPicture struct {
	picture config.Picture
}

// 私聊
func (p *PrivateText) Controls(s any) {
	if strings.Contains(s.(config.Messages).Message, "课表") {
		for _, i2 := range s.(config.Messages).Message {
			//是否为数字
			if unicode.IsNumber(i2) {
				p.class.w, _ = strconv.ParseInt(string(i2), 10, 64)
				fmt.Println(" p.class.w:", p.class.w)
				p.class.GetClass()
			}
		}
		config.SendChan <- p.class.GetClass()
	}

	//return p.chatgpt.GetMessage("你是谁")

	//s.(config.Messages).Message
	//fmt.Println(69)
}

func (p *PrivatePicture) Controls(s any) {
	if strings.Contains(s.(config.Messages).Message, "定时") {
		str := strings.Split(strings.ReplaceAll(s.(config.Messages).Message, "  ", ""), "|")
		p.t.Message = "哈哈哈"
		p.t.Time()
		log.Panicln("str", str)

	}

	if strings.Contains(s.(config.Messages).Message, "课表") {
		i := ""
		for _, i2 := range s.(config.Messages).Message {
			//是否为数字
			if unicode.IsNumber(i2) {
				i += string(i2)
			}
		}
		p.class.w, _ = strconv.ParseInt(i, 10, 64)
		fmt.Println(" p.class.w:", p.class.w)
		p.picture.CreatePicture(p.class.GetClass())
		fmt.Println("生成完成")
	} else {
		p.picture.CreatePicture(s.(config.Messages).Message)
	}
	//消息防撤回
	if s.(config.Messages).Notice_type == "friend_recall" && config.K.Mode.Recall {
		p.m.preventRecall(s.(config.Messages))
		fmt.Println("p.m.c.Message", M.Data.Message)
		p.picture.CreatePicture(M.Data.Message)
	}

	//制作图片
	config.SendChan <- "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
}

// 群聊
func (g *GroupText) Controls(s any) {

}

func (g *GroupPicture) Controls(s any) {

}
