/*
   连接etcd的go client包ཟ在goproxy.cn的镜像里面没有需要git自己拉下来

   Author ： CHR_崔贺然
   Time ： 2019.11.13
   Description ： ETCD Go client
*/
package etcd

// import (
// 	"fmt"
// 	"time"

// 	"github.com/coreos/etcd/clientv3"
// )

// func EtcdInit() {

// 	cli, err := clientv3.New(clientv3.Config{
// 		Endpoints:   []string{"localhost:2379"},
// 		DialTimeout: 5 * time.Second,
// 	})
// 	if err != nil {
// 		fmt.Println("connect failed, err:", err)
// 		return
// 	}
// 	fmt.Println("connect success")
// 	defer cli.Close()
// }
