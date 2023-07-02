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
	MessageDeal(s any) string
}

// 私聊
type Private struct {
	class   class
	chatgpt ChatGpt
	m       manager
	t       timing
	c       collyBaidu
}

// 群聊
type Group struct {
	gh GroupList
	m  manager
}

// 处理其他事件
type Other struct {
	a AutoFriend
}

// 生成文字
type Text struct {
	m mess
}

// 生成图片
type Picture struct {
	m mess
	p config.Picture
}

// 文字
func (t *Text) Controls(s any) {
	if s.(config.Messages).Message_type == "private" {
		t.m = new(Private)
		if t.m.MessageDeal(s) == "" {
			return
		}
		config.SendChan <- t.m.MessageDeal(s)
		return

	} else if s.(config.Messages).Message_type == "group" {
		t.m = new(Group)
		if t.m.MessageDeal(s) == "" {
			return
		}
		config.SendChan <- t.m.MessageDeal(s)
		return
	}
	t.m.MessageDeal(s)
}

// 图片
func (p *Picture) Controls(s any) {
	if s.(config.Messages).Message_type == "private" {
		p.m = new(Private)
		str := p.m.MessageDeal(s)
		if str == "" {
			return
		}
		p.p.CreatePicture(str)
		config.SendChan <- "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
		return
	} else if s.(config.Messages).Message_type == "group" {
		p.m = new(Group)
		str := p.m.MessageDeal(s)
		if str == "" {
			return
		}
		p.p.CreatePicture(str)
		config.SendChan <- "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
		return
	} else if s.(config.Messages).Notice_type != "" {
		//撤回消息
		str := p.m.MessageDeal(s)
		if str == "" {
			return
		}
		p.p.CreatePicture(M.Data.Message)
		config.SendChan <- str + "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
		return
	} else if s.(config.Messages).Post_type == "request" {
		//好友添加
		p.m = new(Other)
		p.m.MessageDeal(s)
	}

	str := p.m.MessageDeal(s)
	if str == "" {
		return
	}
	p.p.CreatePicture(str)
	config.SendChan <- "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
}

// 私聊
func (p *Private) MessageDeal(s any) string {
	st := s.(config.Messages).Message
	if strings.Contains(s.(config.Messages).Message, "定时") {
		//str := strings.Split(strings.ReplaceAll(s.(config.Messages).Message, "  ", ""), "|")
		//要发送的 消息
		p.t.Message = "哈哈哈"
		p.t.Time()
		//log.Panicln("str", str)

	}
	//抢红包
	//爬取百度
	if strings.Contains(s.(config.Messages).Message, "热搜") {
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
				st = string(file)
				k = true
				break
			}
		}
		if k {
			return st
		}
		st = p.c.crawler()
	}

	if strings.Contains(s.(config.Messages).Message, "课表") {
		i := ""
		for _, i2 := range s.(config.Messages).Message {
			//是否为数字
			if unicode.IsNumber(i2) {
				i += string(i2)
			}

		}
		//w为第几周
		p.class.w, _ = strconv.ParseInt(i, 10, 64)
		fmt.Println(" p.class.w:", p.class.w)
		st = p.class.GetClass()

	}

	//消息防撤回
	if s.(config.Messages).Notice_type == "friend_recall" && config.K.Mode.Recall {
		p.m.preventRecall(s.(config.Messages))
		fmt.Println("p.m.c.Message", M.Data.Message)
		st = "[CQ:at," + "qq=" + strconv.FormatInt(s.(config.Messages).User_id, 10) + "]撤回消息" + "\n"
	}

	return st
}

// 群聊
func (g *Group) MessageDeal(s any) string {
	st := s.(config.Messages).Message
	if s.(config.Messages).User_id == 3096407768 {
		if strings.Contains(s.(config.Messages).Message, "禁言") {
			g.gh.receive(s)

		}
		if s.(config.Messages).Notice_type == "group_recall" && s.(config.Messages).Operator_id == s.(config.Messages).User_id && config.K.Mode.Recall {
			g.m.preventRecall(s.(config.Messages))
			fmt.Println("p.m.c.Message", M.Data.Message)
			st = "[CQ:at," + "qq=" + strconv.FormatInt(s.(config.Messages).User_id, 10) + "]撤回消息" + "\n"
		}

	} else {
		st = "您没有权限"
	}

	return st
}

func (o *Other) MessageDeal(s any) string {
	if s.(config.Messages).Request_type == "friend" {
		o.a.auto(s)
	}
	return ""
}

//func (p *Private) Controls(s any) {
//	if strings.Contains(s.(config.Messages).Message, "定时") {
//		str := strings.Split(strings.ReplaceAll(s.(config.Messages).Message, "  ", ""), "|")
//		p.t.Message = "哈哈哈"
//		p.t.Time()
//		log.Panicln("str", str)
//	}
//
//	if strings.Contains(s.(config.Messages).Message, "课表") {
//		i := ""
//		for _, i2 := range s.(config.Messages).Message {
//			//是否为数字
//			if unicode.IsNumber(i2) {
//				i += string(i2)
//			}
//		}
//		//w为第几周
//		p.class.w, _ = strconv.ParseInt(i, 10, 64)
//		fmt.Println(" p.class.w:", p.class.w)
//		p.picture.CreatePicture(p.class.GetClass())
//		fmt.Println("生成完成")
//	} else {
//		p.picture.CreatePicture(s.(config.Messages).Message)
//	}
//	//消息防撤回
//	if s.(config.Messages).Notice_type == "friend_recall" && config.K.Mode.Recall {
//		p.m.preventRecall(s.(config.Messages))
//		log.Println("p.m.c.Message", M.Data.Message)
//		p.picture.CreatePicture(M.Data.Message)
//		config.SendChan <- "[CQ:at," + "qq=" + strconv.FormatInt(s.(config.Messages).User_id, 10) + "]撤回消息" + "\n" + "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
//		return
//	}
//	//热搜
//	if strings.Contains(s.(config.Messages).Message, "热搜") {
//		t := time.Now().Format("2006-01-02")
//		filelist, err := os.ReadDir("./config")
//		if err != nil {
//			log.Panicln("打开文件错误" + err.Error())
//		}
//		for _, v := range filelist {
//			//判断今天是否爬取
//			if strings.Split(v.Name(), ".")[0] == t {
//				file, err := os.ReadFile("./config/" + v.Name())
//				if err != nil {
//					log.Panicln("读取出错" + err.Error())
//				}
//				p.picture.CreatePicture(string(file))
//				break
//			}
//		}
//		//执行爬取p.c.crawler()
//		p.picture.CreatePicture(p.c.crawler())
//	}
//	//制作图片
//	config.SendChan <- "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
//}
