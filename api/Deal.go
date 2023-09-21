package api

import (
	"Go-Bat/config"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type mess interface {
	MessageDeal(s config.Messages, model string)
}

// Private 私聊
type Private struct {
	class   class
	chatGpt ChatGpt
	m       manager
	t       timing
	c       collyBaidu
	d       Ddns
}

// Group 群聊
type Group struct {
	gh GroupList
	m  manager
}

// Other 处理其他事件
type Other struct {
	a AutoFriend
}

// Text 生成文字
type Text struct {
	m mess
}

// Picture 生成图片
type Picture struct {
	m mess
	p config.Picture
}

// 文字
func (t *Text) Controls(s config.Messages) {
	fmt.Println(s.NoticeType, s.MessageType == "group")
	if s.MessageType == "private" || s.NoticeType != "" {
		t.m = new(Private)
		t.m.MessageDeal(s, "t")
	} else if s.MessageType == "group" {
		t.m = new(Group)
		t.m.MessageDeal(s, "t")
	} else {
		t.m = new(Other)
		t.m.MessageDeal(s, "t")
	}
}

// 图片
func (p *Picture) Controls(s config.Messages) {
	fmt.Println("8fjj", s.NoticeType, s.MessageType == "group")
	if s.MessageType == "private" || s.NoticeType == "friend_recall" {
		p.m = new(Private)
		p.m.MessageDeal(s, "p")
	} else if s.MessageType == "group" {
		p.m = new(Group)
		p.m.MessageDeal(s, "p")
	} else if s.PostType == "request" {
		//好友添加
		p.m = new(Other)
		p.m.MessageDeal(s, "p")
	}

	go func() {
		for {
			select {
			case c := <-config.PicterChan:
				log.Println("读取图片生成", c.Message)
				{
					p.p.CreatePicture(c.Message)
					config.SendChan <- config.SendMessage{
						UserId:      c.UserId,
						GroupId:     c.GroupId,
						Message:     "[CQ:image,file=file:///root/GoBatRoot/config/f.png]",
						MessageType: c.MessageType,
						AutoEscape:  false,
					}
				}
			}
		}
	}()

}

// 私聊
func (p *Private) MessageDeal(s config.Messages, m string) {
	message := config.SendMessage{
		UserId:      s.UserId,
		Message:     "",
		MessageType: s.MessageType,
		AutoEscape:  false,
	}

	if strings.Contains(s.Message, "机器人") && s.UserId == config.K.Bat.QQ {
		config.K.Mode.Chatgpt = !config.K.Mode.Chatgpt
		s := ""
		if config.K.Mode.Chatgpt {
			s = "打开"
		} else {
			s = "关闭"
		}
		message.Message = s

	}

	if strings.Contains(s.Message, "更新ip") && s.UserId == config.K.Bat.QQ {
		ok := p.d.Times()
		if ok {
			message.Message = "更新成功"
			return
		}
		message.Message = "更新失败"

	}

	if strings.Contains(s.Message, "定时") {
		//str := strings.Split(strings.ReplaceAll(s.(config.Messages).Message, "  ", ""), "|")
		//要发送的 消息
		p.t.Message = "哈哈哈"
		p.t.Time(s)
		//log.Panicln("str", str)
	}
	//爬取百度
	if strings.Contains(s.Message, "热搜") {
		t := time.Now().Format("2006-01-02")
		k := false
		filelist, err := os.ReadDir("./config")
		if err != nil {
			log.Panicln("打开文件错误" + err.Error())
		}
		for _, v := range filelist {
			//判断今天是否爬取
			if strings.Split(v.Name(), ".")[0] == t {
				file, err := os.ReadFile("./config/" + v.Name())
				if err != nil {
					log.Panicln("读取出错" + err.Error())
				}
				if string(file) != "" {
					message.Message = string(file)
					k = true
					break
				}
			}
		}
		if !k {
			message.Message = p.c.crawler()
		}
	}

	if strings.Contains(s.Message, "课表") {
		if s.UserId != config.K.Bat.QQ {
			message.Message = "您没有权限查看课表"
			return
		}
		i := ""
		for _, i2 := range s.Message {
			//是否为数字
			if unicode.IsNumber(i2) {
				i += string(i2)
			}
		}
		//w为第几周
		p.class.w, _ = strconv.ParseInt(i, 10, 64)
		fmt.Println(" p.class.w:", p.class.w)
		message.Message = p.class.GetClass()
	}
	if strings.Contains(s.Message, "元神启动") {
		go p.class.SetTime()
	}
	//消息防撤回
	if s.NoticeType == "friend_recall" && config.K.Mode.Recall {
		p.m.preventRecall(s)
		fmt.Println("p.m.c.Message", M.Data.Message)
		message.Message = "[CQ:at," + "qq=" + strconv.FormatInt(s.UserId, 10) + "]  撤回消息" + "\n" + M.Data.Message
	}
	if strings.Contains(s.Message, "CQ:face") || strings.Contains(s.Message, "CQ:image") {
		if m == "p" {
			m = "t"
		}
		message.Message = s.Message
	}

	if config.K.Mode.Chatgpt && message.Message == "" && !strings.Contains(s.Message, "[CQ:image,") {
		err := p.chatGpt.Limit(s)
		if err != "" {
			message.Message = err
			return
		}
		message.Message = p.chatGpt.GetMessage(s.Message)
	}

	if m == "t" {
		if message.Message == "" {
			return
		}
		config.SendChan <- message
		return
	}
	if message.Message == "" {
		message.Message = s.Message
	}
	log.Println(message.Message)
	config.PicterChan <- message
}

// 群聊
func (g *Group) MessageDeal(s config.Messages, m string) {
	messages := config.SendMessage{
		GroupId:     s.GroupId,
		Message:     "",
		MessageType: s.MessageType,
		AutoEscape:  false,
	}
	if s.UserId == 3096407768 {
		//禁言
		if strings.Contains(s.Message, "禁言") {
			g.gh.receive(s)
		}
		//防撤回
		if s.NoticeType == "group_recall" && s.OperatorId == s.UserId && config.K.Mode.Recall {
			g.m.preventRecall(s)
			fmt.Println("p.m.c.Message", M.Data.Message)
			messages.Message = "[CQ:at," + "qq=" + strconv.FormatInt(s.UserId, 10) + "]撤回消息" + "\n" + M.Data.Message
		}
		//	撤回消息
		if strings.Contains(s.Message, "撤回") {
			uid := s.MessageId
			s.MessageId, _ = strconv.ParseInt(s.Message[strings.Index(s.Message, "=")+1:strings.Index(s.Message, "]")], 10, 64)
			fmt.Println("s.message", s.MessageId)
			g.m.ban(s)
			s.MessageId = uid
			g.m.ban(s)
		}
	} else {
		messages.Message = "您没有权限"
	}

	if m == "t" {
		if messages.Message == "" {
			return
		}
		config.SendChan <- messages
		return
	}
	if messages.Message == "" {
		return
	}
	config.PicterChan <- messages
}

// 其他
func (o *Other) MessageDeal(s config.Messages, m string) {
	messages := config.SendMessage{
		UserId:      s.UserId,
		GroupId:     s.GroupId,
		Message:     "",
		MessageType: s.MessageType,
		AutoEscape:  false,
	}
	if s.RequestType == "friend" {
		o.a.auto(s)
	}
	if m == "t" {
		config.SendChan <- messages
		return
	}
	config.PicterChan <- messages
}
