/*
  Author ： CHR_崔贺然
  Time ： 2019.11.14
  Description ： rpc server端代码
  rpc framework ： https://github.com/grpc/grpc (gRPC)
				   https://github.com/apache/thrift (Thrift)
				   https://github.com/Tencent/phxrpc (Tencent，还有Tars)
				   https://github.com/brpc/brpc (Baidu)
*/

package rpc

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// RPC is 对外服务端代码
func RPC() {

	// 服务端会调用(用于HTTP服务)
	arith := new(Arith)

	rpc.Register(arith)

	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)

	if err != nil {
		fmt.Println(err.Error())
	}

}
