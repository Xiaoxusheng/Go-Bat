package main

import (
	"Go-Bat/message"
)

func main() {
	//初始化连接
	GoBat := message.NewGoBat()
	//异步启动监听
	go GoBat.Start()
	//启动服务
	GoBat.Serve()

}
