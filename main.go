package main

import (
	"Go-Bat/abstraction"
	"Go-Bat/message"
)

func main() {

	//创建机器人对象
	GoBat := message.NewGoBat()
	//异步启动监听消息
	go GoBat.Start()

	Gobat := new(abstraction.GoBat)
	//设置模式
	Gobat.Tactics()
	//启动服务
	GoBat.Serve(Gobat)

}
