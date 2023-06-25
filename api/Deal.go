package api

import "Go-Bat/config"

type PrivateText struct {
	class   Class
	chatgpt ChatGpt
}

type PrivatePicture struct {
	class   Class
	picture config.Picture
}

type GroupText struct {
}

type GroupPicture struct {
}

func (p *PrivateText) Controls(s any) any {
	//if strings.Contains(s.(config.Messages).Message, "课程表") {
	//	p.class.GetClass()
	//}
	//return p.chatgpt.GetMessage("你是谁")
	return "你好"
	//s.(config.Messages).Message
	//fmt.Println(69)
}

func (p *PrivatePicture) Controls(s any) any {
	p.picture.CreatePicture(s.(config.Messages).Message)
	return "[CQ:image,file=file:///www/Go-Bat/www/wwwroot/GoBat/config/f.png]"
}

func (g *GroupText) Controls(s any) any {
	return ""
}

func (g *GroupPicture) Controls(s any) any {
	return ""

}
