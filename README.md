## GoBat

### 基于go-cqhttp与GO语言实现
https://github-readme-stats.vercel.app/api?username=Xiaoxusheng&theme=blue-green
https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
<img src="https://github-readme-stats.vercel.app/api?username=Xiaoxusheng&theme=blue-green">
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go badge">
<img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/Xiaoxusheng/Go-Bat">
<img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/Xiaoxusheng/Go-Bat">
<img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/Xiaoxusheng/Go-Bat">
<img alt="GitHub contributors" src="https://img.shields.io/github/contributors/Xiaoxusheng/Go-Bat">
### 1 功能如下
- [x] 私聊
  - [x] 天气模式: (t)
  - [x] 聊天模式：(chat 接入chatgpt)
  - [x] 今日新闻： (f)
  - [x] 定时消息推送：(h 例 12-4-QQ号-早早早)
  - [x] 学习通课表提送功能 (w 第三周发送3)
  - [x] 私人消息防撤回
  - [x] 自动同意添加好友
- [x] 群聊
  - [x] 撤回消息
  - [x] 聊天（与chatgpt聊天）
  - [x] 群消息防撤回
  - [x] 禁言群成员，取消禁言 （@群成员 禁言 x 分钟 0代表解除禁言）
  - [x] 消息防撤回功能的开关（可以自己打开防撤回功能）
  - [x] 每日发送消息数（零点自动清零）
  - [x] chatgpt画图功能 （画图 文字描述）
  - [x] 群机器人开关功能（机器人关闭，机器人关闭）

### 2配置 go-cqhttp config.yml

#### 1. 进入配置qq账号密码

         uin: xxxxxxxx  # QQ账号
         password: xxxxxx # 密码为空时使用扫码登录

#### 2. 选择通信方式为 0 3

     - http: # HTTP 通信设置
      address: 0.0.0.0:5000 # HTTP监听地址
      timeout: 5      # 反向 HTTP 超时时间, 单位秒，<5 时将被忽略
      long-polling:   # 长轮询拓展
        enabled: false       # 是否开启
        max-queue-size: 2000 # 消息队列大小，0 表示不限制队列大小，谨慎使用
      middlewares:
        <<: *default # 引用默认中间件
      post:           # 反向HTTP POST地址列表
       # 反向WS设置
      - ws-reverse:
     # 反向WS Universal 地址
     # 注意 设置了此项地址后下面两项将会被忽略
      universal: ws://127.0.0.1:5700
      # 反向WS API 地址
      api: ws://your_websocket_api.server
      # 反向WS Event 地址
      event: ws://your_websocket_event.server
      # 重连间隔 单位毫秒
      reconnect-interval: 3000
      middlewares:
      <<: *default # 引用默认中间件

#### 3.开启服务

         双击go-cqhttp.bat
       ./go-cqhttp  enter运行

#### 4.服务器启动

      node app.js

#### 5.api接口请访问

<https://docs.go-cqhttp.org/>

#### 6.结果



#### 6.说明：

     部署到阿里云或者腾讯云服务器上查询学习通课表无法使用，学习通屏蔽了服务器ip

### 3.配置在config.json中



### 4.声明 练手学习使用，无其他用途