/*
  Author ： CHR_崔贺然
  Time ： 2019.11.14
  Description ： rpc client端代码
*/

package rpc

import (
	"fmt"
	"log"
	"net/rpc"
)

type ClientArgs struct {
	A, B int
}

type ClientQuotient struct {
	Quo, Rem int
}

func RpcClient() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		log.Fatal("dialing", err)
	}
	args := ClientArgs{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
