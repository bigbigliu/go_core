# go-core

golang中使用 `import`引入，使用基本库方法，如果有特殊需求，可以继承并重写类方法实现

```golang
import "github.com/bigbigliu/go_core"
```

## 目录结构
```txt
|-- go-core
├── config
│   ├── conf.go
│   └── settings-dev.yml
├── database
│   ├── mysql
│   │   ├── dbClient.go
│   │   └── models.go
│   └── redis
│       └── redisClient.go
├── logger
│   ├── gin.go
│   ├── gorm.go
│   └── logger.go
├── main.go
├── pkgs
│   ├── error.go
│   ├── getIP.go
│   ├── id.go
│   ├── id_test.go
│   ├── return.go
│   └── storage
│       ├── aliyunOss
│       │   └── upload.go
│       ├── qiniuOss
│       │   └── upload.go
│       └── upyunOss
│           └── upload.go
├── template
│   └── db
│       └── model.go
└── web
    ├── jwt_token
    │   ├── gin.go
    │   └── jwt.go
    └── web_middleware
        ├── cors.go
        ├── ipFilter.go
        ├── ipFilterRedis.go
        ├── requestID.go
        └── timeout.go

```