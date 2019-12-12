package service

import (
	"fmt"
	"net/http"
	"os"
	"project2019/cache"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// 测试基本写法
func RedisGetKey(c *gin.Context) {
	// 获取的时候连接一下
	conn, _ := redis.Dial("tcp", os.Getenv("REDIS_CONN"), redis.DialPassword(os.Getenv("REDIS_PW")))

	// set数据
	_, errSet := conn.Do("SET", "ping", "pong")
	if errSet != nil {
		fmt.Println("redis set failed", errSet)
	}
	fmt.Println("redis set success")
	// redis的get操作
	res, err := conn.Do("GET", "cui")
	if err != nil {
		fmt.Println("testing get failed", err)
		return
	}

	// res 返回一个接口，需要通过断言拿到里面的数据
	switch t := res.(type) {
	case []byte:
		//    fmt.Println(string(t))
		c.String(http.StatusOK, "return:", string(t))
	}

}

// 测试Redis连接池

func RedisPool(c *gin.Context) {
	// 从连接池取一个连接
	conn := cache.POOL.Get() // 从连接池取一个链接
	defer conn.Close()       // 函数运行结束把链接放回连接池

	_, err := conn.Do("SET", "abc", 666)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(conn.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	fmt.Println(r)
	c.String(http.StatusOK, "", r)
	cache.POOL.Close() //关闭连接池

}
