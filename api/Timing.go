package api

import (
	"Go-Bat/config"
	"fmt"
	"time"
)

type Timing struct {
	Message string
	Number  int64
	Private string
}

func (t *Timing) Time() {
	now := time.Now()
	nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 8, now.Minute(), now.Second()+10, 0, now.Location()).Sub(now)
	//nextDay := now.Add(time.Second * 10)
	timer := time.NewTimer(nextDay)
	fmt.Println("开始")
	//异步
	go func() {
		select {
		case <-timer.C:
			//重置
			now = time.Now()
			nextDay = time.Date(now.Year(), now.Month(), now.Day()+1, 8, now.Minute(), now.Second()+10, 0, now.Location()).Sub(now)
			timer = time.NewTimer(nextDay)
			config.SendChan <- t.Message
		}
	}()

}
