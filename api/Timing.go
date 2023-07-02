package api

import (
	"Go-Bat/config"
	"fmt"
	"time"
)

type timing struct {
	Message string
	Number  int64
	Private string
}

func (t *timing) Time() {
	now := time.Now()
	nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, now.Hour(), now.Minute(), now.Second(), 0, now.Location()).Sub(now)
	timer := time.NewTimer(nextDay)
	fmt.Println("开始")
	//异步
	go func() {
		for {
			select {
			case <-timer.C:
				//重置
				now = time.Now()
				nextDay = time.Date(now.Year(), now.Month(), now.Day()+1, now.Hour(), now.Minute(), now.Second(), 0, now.Location()).Sub(now)
				_ = timer.Reset(nextDay)
				fmt.Println(timer.C)
				config.SendChan <- t.Message
			}
		}
	}()

}
