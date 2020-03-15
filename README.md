# Python To Golang

> Gin脚手架

---
# Gin 项目模板

- 配置 Golang 环境
- 配置 goproxy

```
go mod init project2019
go run main.go
export GIN_MODE=release
```

- 在根目录创建.env 文件添加项目环境变量（连接数据库信息，各种密码，token 等）、CORS 跨域中间件、登录认证中间件

```
MYSQL_DSN="username:password!@(hostip)/dbname?charset=utf8mb4&parseTime=True&loc=Local"  # Mysql
REDIS_CONN="hostip:6379"  # Redis
SESSION_SECRET=""  # Token
REDIS_PW="password"  # Redis_Password
KAFKA_CONF="hostip:9092"  # Kafka
ES_CONF="hostip:9200"  # Elasticsearch
```

# 模块内容

- Golang 消费 Mysql
- Golang 消费 Redis
- Golang 消费 Kafka
- Golang 消费 ETCD
- Golang 消费 ElasticSearch
- Golang 七牛云对象存储 API
- Golang Jwt+OAuth2.0认证


