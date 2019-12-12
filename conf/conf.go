/*
  Author ： CHR_崔贺然
  Time ： 2019.11.14
  Description ： 配置中心，所有的配置都在这里：init函数；
  				全部先加载conf里的Init配置信息；
  				main.go(init) -> conf -> models(init.go)/cache(init.go) -> main.go(API)
*/
package conf

import (
	"fmt"
	"log"
	"os"
	"project2019/cache"

	// "project2019/kafka"

	model "project2019/models"

	//"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

func Init() {
	// 连接数据库
	// models.Datebase()
	err_mysql := godotenv.Load()
	if err_mysql != nil {
		log.Fatal("Error loading .env file")
	}
	model.Datebase(os.Getenv("MYSQL_DSN"))
	fmt.Println("mysql conn success")
	// 设置日至级别
	//utils.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接Redis
	// cache.ConnRedis()

	// ་建立redis连接池
	cache.RedisPool()

	// 连接ETCD
	// etcd.EtcdInit()

	// 连接kafka
	// kafka.KafkaInit()

	// 连接Elasticsearch
	// elasticsearch.EsInit()

	// 七牛云认证
	// qiniu.QiniuAuth()
}
