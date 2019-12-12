package cache

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

// 基本写法，实际业务用连接池（init.go）中
func ConnRedis() {
	// 在环境变量中获取redis的配置，注意最后一项的密码写法
	c, err := redis.Dial("tcp", os.Getenv("REDIS_CONN"), redis.DialPassword(os.Getenv("REDIS_PW")))

	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}
	fmt.Println("redis conn success")

	defer c.Close()
	// fmt.Println(POOL)
}
