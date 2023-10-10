package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/jordan-wright/email"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type Ddns struct{}

type T struct {
	Code int `json:"code"`
	Data struct {
		Myip     string `json:"myip"`
		Location string `json:"location"`
		Country  string `json:"country"`
		Local    string `json:"local"`
		Ver4     string `json:"ver4"`
		Ver6     string `json:"ver6"`
		Count4   int    `json:"count4"`
		Count6   int    `json:"count6"`
	} `json:"data"`
}

//此脚本用来动态解析dns,每天请求一次,记录ipv6地址,当发现ipv6变化时候,主动修改阿里云dns解析

func GetIpv6() (string, bool) {
	req, err := http.NewRequest("GET", "https://v6.ip.zxinc.org/info.php?type=json", nil)
	if err != nil {
		log.Println(err)
		return "", false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.1938.62")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", false

	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", false

	}
	t := new(T)
	err = json.Unmarshal(res, t)
	if err != nil {
		log.Println("json解析错误")
		return "", false
	}
	log.Println("Ipv6获取成功", t.Data.Myip)
	return t.Data.Myip, true
}

func GetIpv4() (string, bool) {
	req, err := http.NewRequest("GET", "https://v4.ip.zxinc.org/info.php?type=json", nil)
	if err != nil {
		log.Println(err)
		return "", false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.1938.62")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取失败" + err.Error())
		return "", false

	}
	v := new(T)
	err = json.Unmarshal(res, v)
	if err != nil {
		log.Println("json解析错误")
		return "", false
	}
	log.Println("ipv4获取成功", v.Data.Myip)
	return v.Data.Myip, true

}

func Set(v4, v6 string) bool {
	config := sdk.NewConfig()

	// Please ensure that the environment variables ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set.
	credential := credentials.NewAccessKeyCredential("", "")
	/* use STS Token
	credential := credentials.NewStsTokenCredential(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"), os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN"))
	*/
	client, err := alidns.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		log.Println(err)
	}

	request := alidns.CreateUpdateDomainRecordRequest()

	request.Scheme = "https"
	request.Type = "AAAA"
	request.Value = v6
	request.RR = "@"
	request.RecordId = "848798676399225856"
	request.Lang = "en"
	request.UserClientIp = v4
	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)

	return response.IsSuccess()
}

func SendEmail(v6, v4 string) {
	e := email.NewEmail()
	//发送者
	e.From = "服务器IPV6IPV4地址<2673893724@qq.com>"
	//接收者
	e.To = []string{"3096407768@qq.com"}
	//主题
	e.Subject = "IP地址"
	//html
	e.HTML = []byte("<!DOCTYPE html>\n<html>\n<head>\n    <title>IPv6 & IPv4 Address</title>\n    <style>\n        body {\n            font-family: Arial, sans-serif;\n         " +
		"   background-color: #f7f7f7;\n            margin: 0;\n            padding: 20px;\n        }\n        \n        h1 {\n            text-align: center;\n           " +
		" color: #333;\n        }\n        \n        .container {\n            max-width: 400px;\n            margin: 0 auto;\n            background-color: #fff;\n    " +
		"        border-radius: 5px;\n            box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);\n            padding: 20px;\n        }\n        \n        .address {\n          " +
		"  margin-bottom: 10px;\n        }\n        \n        .label {\n            font-weight: bold;\n        }\n        \n        .ipv6 {\n            font-family: monospace;\n        " +
		"    color: #333;\n            background-color: #f7f7f7;\n            padding: 5px;\n            border-radius: 3px;\n        }\n        \n        .ipv4 {\n            font-family: monospace;\n   " +
		"         color: #333;\n            background-color: #f7f7f7;\n            padding: 5px;\n            border-radius: 3px;\n        }\n    </style>\n</head>\n<body>\n    <div class=\"container\">\n  " +
		"      <h1>IPv6 & IPv4 Address</h1>\n        \n        <div class=\"address\">\n            <span class=\"label\">IPv6:</span>\n          " +
		" <span class=\"ipv6\">" + v6 + "</span>\n       " +
		" </div>\n        \n        <div class=\"address\">\n           " +
		" <span class=\"label\">IPv4:</span>\n          " +
		"  <span class=\"ipv4\">" + v4 + "</span>\n      " +
		"  </div>\n    </div>\n</body>\n</html> ")

	err := e.SendWithStartTLS("smtp.qq.com:587", smtp.PlainAuth("", "", "", "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.gmail.com:465"})
	if err != nil {
		log.Println("stmp:", err)

	}
	log.Println("发送成功！")
}

func (d *Ddns) Times() bool {
	file, err := os.OpenFile("ip.txt", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Println("打开失败！" + err.Error())
		return false
	}
	b := make([]byte, 1024)

	n, err := file.Read(b)
	if err != nil {
		if err == io.EOF {
			file, err := os.OpenFile("ip.txt", os.O_CREATE|os.O_WRONLY, 0755)
			_, err = file.Write([]byte("122"))
			if err != nil {
				log.Println("文件写入失败", err)
				return false
			}
		}
		log.Println("文件读取失败" + err.Error())
		return false
	}
	defer file.Close()

	v6, k := GetIpv6()
	v4, ok := GetIpv4()
	if k && ok {
		SendEmail(v6, v4)
	}
	fmt.Println(string(b[:n]) == v6, string(b[:n]), v6)
	if string(b[:n]) != v6 {
		file, err := os.OpenFile("ip.txt", os.O_WRONLY, 0755)
		_, err = file.Write([]byte(v6))
		if err != nil {
			log.Println("文件写入失败", err)
			return false
		}
		file.Close()
		if ok && k {
			f := Set(v4, v6)
			if f {
				log.Println("设置成功")
				return true
			}
			fmt.Println("设置失败！")
			return false
		}
		log.Println("获取ipv4或ipv6失败")
		return false
	}
	log.Println("ipv6地址没有改变")
	return true
}
