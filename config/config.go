package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Messages struct {
	PostType    string  `json:"post_type,omitempty"`
	MessageType string  `json:"message_type,omitempty"`
	Time        int64   `json:"time,omitempty"`
	SelfId      int64   `json:"self_id,omitempty"`
	SubType     string  `json:"sub_type,omitempty"`
	MessageId   int64   `json:"message_id,omitempty"`
	UserId      int64   `json:"user_id,omitempty"`
	TargetId    int64   `json:"target_id,omitempty"`
	Message     string  `json:"message,omitempty"`
	Sender      *Sender `json:"sender,omitempty"`
	NoticeType  string  `json:"notice_type,omitempty"`
	OperatorId  int64   `json:"operator_id,omitempty"`
	GroupId     int64   `json:"group_id,omitempty"`
	GroupName   string  `json:"group_name,omitempty"`
	GroupMemo   string  `json:"group_memo,omitempty"`
	RequestType string  `json:"request_type,omitempty"`
	Flag        string  `json:"flag,omitempty"`
	Remark      string  `json:"remark,omitempty"`
}

/*
{"post_type":"message","message_type":"group","time":1688259987,"self_id":2673893724,"sub_type":"normal","message_seq":2912,"user_id":3096407768,"anonymous":null,"font":0,"group_id":682671449,"message":"[CQ:at,qq=2044139249] 禁言10分钟","raw_message":"[CQ:at,qq=2044139249] 禁言10分钟",
"sender":{"age":0,"area":"","card":"","level":"","nickname":"Ra","role":"member","sex":"unknown","title":"","user_id":3096407768},"message_id":1826682242}*/

type Sender struct {
	Age    int64  `json:"age,omitempty"`
	Sex    string `json:"sex,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}

type SendMessage struct {
	UserId      int64  `json:"user_id"`
	GroupId     int64  `json:"group_id"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	AutoEscape  bool   `json:"auto_escape"`
}

// 处理完数据的管道
var SendChan = make(chan SendMessage, 100)

// 生成图片管道
var PicterChan = make(chan SendMessage, 10)

// 接收消息管道
var MessageChan = make(chan Messages, 100)

// 要已读数据管道
var ReadChan = make(chan int64, 10)

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
		Modes   string
		Bat     bool
		Recall  bool
		Chatgpt bool
		Key     string
	}
	ChaoXing struct {
		Name     string
		Password string
	}
	Bat struct {
		QQ int64
	}
}

type Class struct {
	Result int         `json:"result"`
	Msg    interface{} `json:"msg"`
	Data   struct {
		Curriculum struct {
			Fid                   int         `json:"fid"`
			RealCurrentWeek       int         `json:"realCurrentWeek"`
			CourseTypeName        interface{} `json:"courseTypeName"`
			EarlyMorningTime      interface{} `json:"earlyMorningTime"`
			FirstWeekDateReal     int64       `json:"firstWeekDateReal"`
			UserFid               int         `json:"userFid"`
			Uuid                  string      `json:"uuid"`
			SectionTime           interface{} `json:"sectionTime"`
			Puid                  int         `json:"puid"`
			EarlyMorningSection   interface{} `json:"earlyMorningSection"`
			LessonLength          int         `json:"lessonLength"`
			CurriculumCount       int         `json:"curriculumCount"`
			MorningTime           interface{} `json:"morningTime"`
			SchoolYear            string      `json:"schoolYear"`
			CurrentWeek           int         `json:"currentWeek"`
			Id                    int         `json:"id"`
			IsHasEduLesson        int         `json:"isHasEduLesson"`
			AfternoonTime         interface{} `json:"afternoonTime"`
			ExistMaxLength        int         `json:"existMaxLength"`
			MorningSection        interface{} `json:"morningSection"`
			GetLessonFromCache    int         `json:"getLessonFromCache"`
			MaxWeek               int         `json:"maxWeek"`
			UpdateTime            int64       `json:"updateTime"`
			Sort                  int         `json:"sort"`
			UserName              string      `json:"userName"`
			FirstWeekDate         int64       `json:"firstWeekDate"`
			InsertTime            int64       `json:"insertTime"`
			BreakTime             interface{} `json:"breakTime"`
			EveningSection        interface{} `json:"eveningSection"`
			AfternoonSection      interface{} `json:"afternoonSection"`
			EveningTime           interface{} `json:"eveningTime"`
			Semester              int         `json:"semester"`
			MaxLength             int         `json:"maxLength"`
			LessonTimeConfigArray []string    `json:"lessonTimeConfigArray"`
		} `json:"curriculum"`
		LessonArray []struct {
			Fid                int         `json:"fid"`
			Role               int         `json:"role"`
			Weeks              string      `json:"weeks"`
			NoteCid            interface{} `json:"noteCid"`
			MeetCode           interface{} `json:"meetCode"`
			OnlineLocation     interface{} `json:"onlineLocation"`
			MoocMirrorDomain   interface{} `json:"moocMirrorDomain"`
			BeginNumber        int         `json:"beginNumber"`
			LessonConfigId     string      `json:"lessonConfigId"`
			ClassName          string      `json:"className"`
			ClassId            int         `json:"classId"`
			DayOfWeek          int         `json:"dayOfWeek"`
			CourseId           int         `json:"courseId"`
			TeachPlanName      interface{} `json:"teachPlanName"`
			IsMirror           int         `json:"isMirror"`
			TeacherName        string      `json:"teacherName"`
			Length             int         `json:"length"`
			WeekType           int         `json:"weekType"`
			LessonId           string      `json:"lessonId"`
			UnitNoteUrl        interface{} `json:"unitNoteUrl"`
			ShowTeachPlan      interface{} `json:"showTeachPlan"`
			UpdateTime         int64       `json:"updateTime"`
			DailyLessonNoteCid interface{} `json:"dailyLessonNoteCid"`
			TeacherNo          string      `json:"teacherNo"`
			LessonConfigUuid   string      `json:"lessonConfigUuid"`
			TeachPlanId        interface{} `json:"teachPlanId"`
			CourseNoteCid      interface{} `json:"courseNoteCid"`
			CourseNo           string      `json:"courseNo"`
			Name               string      `json:"name"`
			ClassNo            string      `json:"classNo"`
			EducationalNo      interface{} `json:"educationalNo"`
			PersonId           interface{} `json:"personId"`
			Location           string      `json:"location"`
			PptObjectId        interface{} `json:"pptObjectId"`
		} `json:"lessonArray"`
	} `json:"data"`
}

type Cj struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

var K = Config{}

func init() {
	// 创建一个 Logger 对象，同时输出到文件和控制台
	fmt.Println("     __________  ____  ___  ______")
	fmt.Println("    / ____/ __ \\/ __ )/   |/_  __/")
	fmt.Println("   / / __/ / / / __  / /| | / /   ")
	fmt.Println("  / /_/ / /_/ / /_/ / ___ |/ /    ")
	fmt.Println("  \\____/\\____/_____/_/  |_/_/     ")
	fmt.Println("                                  ")
	fmt.Println("[github地址]:https://github.com/Xiaoxusheng/Go-Bat")
	log.SetPrefix("[Go-Bat]-------")
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	err = viper.Unmarshal(&K)
	viper.SetDefault("server.port", 5000)
	viper.SetDefault("server.ws", 5700)
	viper.SetDefault("redis.addr", "127.0.0.1:6379")
	viper.SetDefault("mode.modes", "T")
	viper.SetDefault("mode.key", "")
	viper.SetDefault("redis.poolSize", 1000)
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.password", "admin123")
	viper.SetDefault("mode.bat", false)
	viper.SetDefault("mode.recall", true)
	viper.SetDefault("mode.chatgpt", true)
	viper.SetDefault("chaoXing.name", "19888340365")
	viper.SetDefault("chaoXing.password", "lei125608")
	viper.SetDefault("bat.qq", 3096407768)
	log.Println(K)
	if err != nil {
		log.Println("初始化失败")
		return
	}

}
