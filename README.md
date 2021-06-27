# go init
## 目录结构
```
.
├── README.md
├── application
│   ├── controllers
│   │   ├── agentaction
│   │   ├── test
│   │   └── users
│   ├── logic
│   └── models
│       └── user.go
├── conf
│   ├── systeminfo.ini
│   ├── systeminfo.json
│   └── systeminfo_bak.ini
├── dao
│   ├── mysql
│   │   └── ginmysql.go
│   └── redis
│       └── ginredis.go
├── go.mod
├── go.sum
├── logs
│   ├── 2021-06-26.log
│   ├── 2021-06-27.log
│   └── test.log
├── main.go
├── pkg
├── routers
│   └── routers.go
├── settings
│   └── setting.go
├── src
├── statics
│   └── images
│       └── qrcode
└── tools
    ├── ginlog.go
    └── qrcode.go
```