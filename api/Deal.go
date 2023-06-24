package api

type PrivateText struct {
	class   Class
	chatgpt ChatGpt
}

type PrivatePicture struct {
	class Class
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

func (g *GroupText) Controls(s any) any {
	return ""
}

func (g *GroupPicture) Controls(s any) any {
	return ""

}
