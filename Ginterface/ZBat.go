package Ginterface

import (
	"Go-Bat/config"
)

// 处理事件接口
type Bat interface {
	Controls(s config.Messages)
}

type Mess interface {
	MessageDeal(s config.Messages, model string)
}

type Times interface {
	SetTime()
}
