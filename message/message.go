package message

import (
	"Go-Bat/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type GoBat struct {
	name    string
	version float64
	time    string
}

var Mess config.Messages

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// NewGoBat 创建对象
func NewGoBat() *GoBat {
	return &GoBat{name: "Go-Bat", version: 0.4, time: time.Now().Format("2006-01-02 15:04:05")}
}

// Websocket  异步监听消息，通过chan传递消息
func (b *GoBat) Websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	for {
		err := conn.ReadJSON(&Mess)
		if err != nil {
			continue
		}
		log.Println("解析mess", Mess.PostType)
		if Mess.PostType == "meta_event" {
			log.Println("chan还剩", 100-len(config.MessageChan))
			Mess = config.Messages{}
			continue
		}
		config.MessageChan <- Mess
		Mess = config.Messages{}
	}

}

// Start 开始监听
func (b *GoBat) Start() {
	b.Err()
	//启动websocket服务
	go b.Service()
	//	启动读协程
	go b.ReadMessage()
	//	启动写协程
	go b.WriteMessage()
	//已读消息
	go b.Read()
	select {}
}

// ReadMessage 读取
func (b *GoBat) ReadMessage() {
	//读取管道消息
	ctx := NewZBat()
	ctx.Tactics()
	for {
		select {
		case c := <-config.MessageChan:
			//已读消息
			config.ReadChan <- c.MessageId
			ctx.Deal(c)
			// 如果MessageChan成功读到数据，则进行该case处理语句
			log.Println("收到Mess", c, "\n", "还剩", 100-len(config.MessageChan), c.MessageId)
		}
	}
}

// WriteMessage 写
func (b *GoBat) WriteMessage() {
	for {
		select {
		case c := <-config.SendChan:
			if c.UserId == 0 && c.GroupId == 0 {
				c.UserId = 3096407768
				c.MessageType = "private"
			}
			fmt.Println("读取到数据", c.Message)

			go func() {
				fmt.Println("发送", c)
				marshal, err := json.Marshal(c)
				if err != nil {
					log.Println(err)
					panic(err)
				}
				resp, err := http.Post("http://127.0.0.1:"+strconv.Itoa(config.K.Server.Port)+"/send_msg", "application/json", bytes.NewBuffer(marshal))
				if err != nil {
					log.Println(err)
				}
				resp.Body.Close()
				log.Println("发送成功")
			}()
			ctx := context.Background()
			if config.Rdb.Exists(ctx, "num").Val() == 0 {
				t1 := time.Now()
				t2 := time.Date(t1.Year(), t1.Month(), t1.Day()+1, 0, 0, 0, 0, t1.Location())
				fmt.Println(t2.Sub(t1))
				_, err := config.Rdb.Set(ctx, "num", 1, t2.Sub(t1)).Result()
				if err != nil {
					log.Println(err)
				}
			}
			result, err := config.Rdb.Incr(ctx, "num").Result()
			if err != nil {
				log.Println(result, err)
			}

		}
	}
}

// Read 已读消息
func (b *GoBat) Read() {
	for {
		select {
		case c := <-config.ReadChan:
			_, err := http.Get("http://127.0.0.1:5000/get_forward_msg?message_id=" + strconv.FormatInt(c, 10))
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("已读成功")
		}
	}
}

func (b *GoBat) Err() {
	defer func() {
		if err := recover(); err != nil {
			config.SendChan <- config.SendMessage{
				UserId:     3096407768,
				Message:    err.(string),
				AutoEscape: false,
			}
		}
	}()
}

// Service 服务
func (b *GoBat) Service() {
	//记录日志
	logFile, err := os.OpenFile("GoBat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(logFile, os.Stderr))
	log.Printf("%v v%v %v", b.name, b.version, "机器人启动")
	http.HandleFunc("/", b.Websocket)
	err = http.ListenAndServe(":"+strconv.Itoa(config.K.Server.Ws), nil)
	if err != nil {
		log.Println(err)
		return
	}
}
