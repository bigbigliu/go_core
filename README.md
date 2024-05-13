# go-core

觅知网，自研使用的go框架包

golang中使用`import`引入，使用基本库方法，如果有特殊需求，可以继承并重写类方法实现

```
import "github.com/bigbigliu/go_core"
```

## 目录结构

```
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
│   ├── exported.go
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

## 检测ip限流中间件
1. 引入包 
    "github.com/bigbigliu/go_core/web/web_middleware"
2. 路由加载中间件
    router.Use(web_middleware.IPFilterMiddleware(10, time.Duration(5)*time.Second))  

3. 说明：     
    不同项目根据实际情况传递相关参数：   
        参数1  可以访问的次数        
        参数2  时间范围   
    例如上面的参数  (10 ，time.Duration(5)*time.Second)      
    5秒之内只能访问10次，大于10次抛出异常 {"code":"-1","msg":"Too many requests"}        
    5秒后方可继续访问