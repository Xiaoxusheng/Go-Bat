package Ginterface

import (
	"net/http"
)

type GoBat interface {
	//启动
	Start()
	//监听解析消息
	Websocket(w http.ResponseWriter, r *http.Request)
	//写消息
	WriteMessage()
	//读取消息
	ReadMessage()
	//	已读消息
	Read()
	//	监听全局错误
	Err()
}
