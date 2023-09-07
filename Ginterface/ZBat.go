package Ginterface

import "Go-Bat/config"

// 处理事件接口
// 抽象接口
type Bat interface {
	Controls(s config.Messages)
}

type Times interface {
	SetTime()
}
