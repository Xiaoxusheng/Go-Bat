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

type PrivateText struct {
	class   Class
	chatgpt ChatGpt
	m       Manager
	t       Timing
	c       CollyBaidu
}

type PrivatePicture struct {
	class   Class
	picture config.Picture
	m       Manager
	t       Timing
	c       CollyBaidu
}

type GroupText struct {
}

type GroupPicture struct {
	picture config.Picture
}

// 私聊
func (p *PrivateText) Controls(s any) {
	if strings.Contains(s.(config.Messages).Message, "定时") {
		str := strings.Split(strings.ReplaceAll(s.(config.Messages).Message, "  ", ""), "|")
		//要发送的 消息
		p.t.Message = "哈哈哈"
		p.t.Time()
		log.Panicln("str", str)
		return
	}
	//抢红包

	//爬取百度
	if strings.Contains(s.(config.Messages).Message, "热搜") {
		t := time.Now().Format("2006-01-02")
		filelist, err := os.ReadDir("./config")
		if err != nil {
			log.Panicln("打开文件错误" + err.Error())
		}
		for _, v := range filelist {
			//判断今天是否爬取
			if strings.Split(v.Name(), ".")[0] == t {
				file, err := os.ReadFile("./config" + v.Name())
				if err != nil {
					log.Panicln("读取出错" + err.Error())
				}
				config.SendChan <- string(file)
				break
			}
		}
		config.SendChan <- p.c.crawler()
	}

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

	//消息防撤回
	if s.(config.Messages).Notice_type == "friend_recall" && config.K.Mode.Recall {
		p.m.preventRecall(s.(config.Messages))
		fmt.Println("p.m.c.Message", M.Data.Message)
		config.SendChan <- "[CQ:at," + "qq=" + strconv.FormatInt(s.(config.Messages).User_id, 10) + "]撤回消息" + "\n" + M.Data.Message_type
		return
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
		//w为第几周
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
		log.Println("p.m.c.Message", M.Data.Message)
		p.picture.CreatePicture(M.Data.Message)
		config.SendChan <- "[CQ:at," + "qq=" + strconv.FormatInt(s.(config.Messages).User_id, 10) + "]撤回消息" + "\n" + "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
		return
	}
	//热搜
	if strings.Contains(s.(config.Messages).Message, "热搜") {
		t := time.Now().Format("2006-01-02")
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
				p.picture.CreatePicture(string(file))
				break
			}
		}
		//执行爬取p.c.crawler()
		p.picture.CreatePicture(p.c.crawler())
	}
	//制作图片
	config.SendChan <- "[CQ:image,file=file:///www/Go-Bat/config/f.png]"
}

// 群聊
func (g *GroupText) Controls(s any) {

}

func (g *GroupPicture) Controls(s any) {

}
