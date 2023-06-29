package message

import (
	"Go-Bat/abstraction"
	"Go-Bat/api"
	"Go-Bat/config"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Message interface {

	// 接收消息
	websocket(w http.ResponseWriter, r *http.Request)
	//	开始监听消息
	Start()
	//	已读消息
	read()
	// 服务
	Serve()
	//

}

type GoBat struct {
	name    string
	version float64
	time    string
}

var once sync.Once
var bat *GoBat
var Mess config.Messages
var MessageChan = make(chan config.Messages, 100)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// 创建对象
func NewGoBat() *GoBat {
	once.Do(func() {
		bat = &GoBat{name: "Go-Bat", version: 0.2, time: time.Now().Format("2006-01-02 15:04:05")}
	})
	return bat
}

// websocket异步监听消息，通过chan传递消息
func (b *GoBat) websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	for {
		err := conn.ReadJSON(&Mess)
		if err != nil {
			continue
		}
		fmt.Println("解析mess", Mess)
		MessageChan <- Mess
		Mess = config.Messages{}
		log.Println("送到通道", "chan还剩", 100-len(MessageChan))

	}

}

// Start 开始监听
func (b *GoBat) Start() {
	logFile, err := os.OpenFile("GoBat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	// 创建一个 Logger 对象，同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(logFile, os.Stderr))
	log.Printf("[INFO]: %v  v%v  %v", b.name, b.version, "机器人启动")
	http.HandleFunc("/", b.websocket)
	err = http.ListenAndServe(":"+strconv.Itoa(config.K.Server.Ws), nil)
	if err != nil {
		log.Panicln(err)
	}
}

// 接收消息
func (b *GoBat) Serve() {
	Gobat := new(abstraction.GoBat)
	Gobat.SetStrategy(new(api.PrivatePicture))
	for {
		select {
		case c := <-MessageChan:
			// 如果MessageChan成功读到数据，则进行该case处理语句
			log.Println("收到Mess", c, "\n", "还剩", 100-len(MessageChan))
			if c.Message_type != "" {
				b.read()
				Gobat.Deal(&c)
			} else if c.Notice_type != "" {
				b.read()
				Gobat.Deal(&c)
			}
		default:
			// 如果上面都没有成功，则进入default处理流程
			continue
		}
	}
}

// 已读消息
func (b *GoBat) read() {
	_, err := http.Get("http://127.0.0.1:5000/get_forward_msg?message_id=" + strconv.FormatInt(Mess.Message_id, 10))
	if err != nil {
		log.Panicln(err)
		return
	}
}

func (b *GoBat) GetName() string {
	return b.name
}

func (b *GoBat) GetVersion() float64 {
	return b.version
}
