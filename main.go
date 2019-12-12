/*
  Author ： CHR_崔贺然
  Time ： 2019.11.14
  Description ： 项目的入口函数，所有的开始都从main函数里面开始
*/

package main

import (
	"project2019/api"
	"project2019/conf"
)

func main() {

	// 初始化配置 DB,Cache 等
	conf.Init()

	// 装载路由
	res := api.Api()

	// res 要返回一个引擎才能Run()
	_ = res.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
