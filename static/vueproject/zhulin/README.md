<<<<<<< HEAD
# zhulin

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Run your unit tests
```
npm run test:unit
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
=======
# Pythoner 学习 Golang 写的不好请多多指正

> 之前一直从事 Django 框架开发，后来接触到到 Gin，iris，Sanic，Flask 等轻量级框架
> 因为 Gin 性能过于优秀（压测结果我放在了个人博客的文章里），So 决定主攻 Gin 框架

# Gin 项目模板，持续更新

1. 配置 Golang 环境
2. 配置 goproxy

```
go mod init project2019
```

```
go run main.go
```

```
export GIN_MODE=release
```

> 在根目录创建.env 文件添加项目环境变量（连接数据库信息，各种密码，token 等）、CORS 跨域中间件、登录认证中间件

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
- AND SOME

# 未来 ADD 计划

- Golang 实现经典算法
- Golang 实现经典设计模式
- 前端 Vue 项目
- Golang 客户端访问 Python Server，Python 中实现简单机器学习，大数据等内容
- Golang 和 Python 的爬虫
- Hyperledger Fabric 联盟链 Golang 接口实现，以及链码开发模板
- 一些不错的 Python、Golang 的小项目（脚本）
>>>>>>> e75d243920cc95f4555b2b80f00092c98a824044
