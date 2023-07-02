package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

type Messages struct {
	Post_type    string  `json:"post_type,omitempty"`
	Message_type string  `json:"message_type,omitempty"`
	Time         int64   `json:"time,omitempty"`
	Self_id      int64   `json:"self_id,omitempty"`
	Sub_type     string  `json:"sub_type,omitempty"`
	Message_id   int64   `json:"message_id,omitempty"`
	User_id      int64   `json:"user_id,omitempty"`
	Target_id    int64   `json:"target_id,omitempty"`
	Message      string  `json:"message,omitempty"`
	Sender       *Sender `json:"sender,omitempty"`
	Notice_type  string  `json:"notice_type,omitempty"`
	Operator_id  int64   `json:"operator_id,omitempty"`
	Group_id     int64   `json:"group_id,omitempty"`
	Group_name   string  `json:"group_name,omitempty"`
	Group_memo   string  `json:"group_memo,omitempty"`
}

/*
{"post_type":"message","message_type":"group","time":1688259987,"self_id":2673893724,"sub_type":"normal","message_seq":2912,"user_id":3096407768,"anonymous":null,"font":0,"group_id":682671449,"message":"[CQ:at,qq=2044139249] 禁言10分钟","raw_message":"[CQ:at,qq=2044139249] 禁言10分钟",
"sender":{"age":0,"area":"","card":"","level":"","nickname":"Ra","role":"member","sex":"unknown","title":"","user_id":3096407768},"message_id":1826682242}*/

type Sender struct {
	Age     int64  `json:"age,omitempty"`
	Sex     string `json:"sex,omitempty"`
	User_id int64  `json:"user_id,omitempty"`
}

// 处理完数据的管道
var SendChan = make(chan string, 100)

type Config struct {
	Server struct {
		Addr string
		Ws   int
		Port int
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
		PoolSize int
	}
	Mode struct {
		Mode   string
		Bat    bool
		Recall bool
	}
	ChaoXing struct {
		Name     string
		Password string
	}
}

var K = Config{}

func init() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
	viper.SetDefault("server.port", 5000)
	viper.SetDefault("server.ws", 5700)
	viper.SetDefault("redis.addr", "116.198.44.154:6379")
	viper.SetDefault("mode.mode", "T")
	viper.SetDefault("redis.poolSize", 1000)
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.password", "admin")
	viper.SetDefault("mode.bat", false)
	viper.SetDefault("mode.recall", false)
	viper.SetDefault("chaoXing.name", "19888340365")
	viper.SetDefault("chaoXing.password", "admin123")

	err = viper.Unmarshal(&K, func(config *mapstructure.DecoderConfig) {

	})
	fmt.Println(K)
	if err != nil {
		return
	}

}
