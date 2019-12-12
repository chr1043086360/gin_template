package cache

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

// redisPool的单例
// 首字母必须大写才能被外界访问
var POOL *redis.Pool

// 建立Redis连接池
func RedisPool() {
	POOL = &redis.Pool{ // 实例化一个连接池
		MaxIdle:     16,  // 初始化连接数量
		MaxActive:   0,   // 不确定可以用0，表示自定义
		IdleTimeout: 300, // 300秒不使用自动关闭
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS_CONN"), redis.DialPassword(os.Getenv("REDIS_PW")))
		},
	}
	fmt.Println("redis conn success")

}
