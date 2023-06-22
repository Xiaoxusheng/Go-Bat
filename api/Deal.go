package api

import (
	"Go-Bat/config"
	"strings"
)

type PrivateText struct {
	class Class
}

type PrivatePicture struct {
	class Class
}

type GroupText struct {
}

type GroupPicture struct {
}

func (p *PrivateText) Controls(s any) {
	if strings.Contains(s.(config.Messages).Message, "课程表") {
		p.class.GetClass()
	}
}
