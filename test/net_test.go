package test

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"testing"
)

// 服务端
func Test_net(t *testing.T) {
	f, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")

	tcp, err := net.ListenTCP("tcp", f)

	if err != nil {
		log.Panicln(err)
	}
	//data := make([]byte, 1024)
	log.Println("serve is start")
	for {
		acceptTCP, err := tcp.AcceptTCP()
		if err != nil {
			return
		}
		readString, err := bufio.NewReader(acceptTCP).ReadString('\n')
		if err != nil {
			log.Panicln(err)
		}
		_, err = acceptTCP.Write([]byte(readString))
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println("客户端消息", readString)
		fmt.Println("阻塞...")
		for {
			readString, err := bufio.NewReader(acceptTCP).ReadString('\n')
			if err != nil {
				log.Panicln(err)
			}
			fmt.Println("客户端消息", readString)

			//_, err = acceptTCP.Write([]byte(readString))
		}

	}

}

// 客户端
func TestCreateConnect(t *testing.T) {
	f, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	tcp, err := net.DialTCP("tcp", nil, f)
	if err != nil {
		log.Panicln(err)
	}
	for {
		data := make([]byte, 1024)
		_, err := tcp.Write([]byte("hello,world\n"))
		if err != nil {
			log.Panicln(err)
		}
		size, err := tcp.Read(data)
		if err != nil {
			log.Panicln(err)
		}

		log.Println(string(data[:size]) == "hello,world\n", string(data[:size]))

	}

}
