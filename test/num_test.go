package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
	"unicode"
)

type MyIntSlice []int

func (s MyIntSlice) Len() int {
	return len(s)
}

func (s MyIntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s MyIntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Test_num(t *testing.T) {
	str := "idji id1"
	for _, i2 := range str {
		if unicode.IsNumber(i2) {
			fmt.Println(string(i2))
		}
	}
	now := time.Now()
	nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 8, now.Minute(), now.Second()+10, 0, now.Location()).Sub(now)

	fmt.Println(nextDay, time.Since(time.Date(now.Year(), now.Month(), now.Day()+1, 8, now.Minute(), now.Second()+10, 0, now.Location())))

	file, err := os.OpenFile("../config/"+time.Now().Format("2006-01-02")+".txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Panicln("创建hot.txt出错" + err.Error())
	}
	defer file.Close()
	_, err = file.Write([]byte(str))
	if err != nil {
		log.Panicln("写入文件出错" + err.Error())
	}
	//tf:=time.Now().Format("2006-01-02")
	filelist, err := os.ReadDir("../config")
	if err != nil {
		log.Panicln("打开文件错误" + err.Error())
	}
	for _, v := range filelist {
		//fmt.Println("name:", v.Name(), "type", strings.Replace(filepath.Ext(v.Name()), ".", "", 1), v.IsDir())
		fmt.Println(strings.Split(v.Name(), ".")[0] == time.Now().Format("2006-01-02"))
	}
	//open, err := os.Open("../config/t.png")
	//if err != nil {
	//	log.Panicln(err)
	//}
	//w := bufio.NewWriter(open)
	//fmt.Println(w.Size())
	data := MyIntSlice{9, 4, 7, 2, 1}
	sort.Sort(data)
	fmt.Println(data)
	c := context.TODO()
	f := context.WithValue(c, "q", "w")
	fmt.Println(f.Value("q"))
}
