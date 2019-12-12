package rpc

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type JsonArgs struct {
	A, B int
}

type JsonQuotient struct {
	Quo, Rem int
}

type JsonArith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) JsonDivide(args *JsonArgs, quo *JsonQuotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func JsonRpc() {

	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		/*
			ServeConn在单个连接上执行DefaultServer。ServeConn会阻塞，服务该连接直到客户端挂起。调用者一般应另开线程调用本函数："go serveConn(conn)"。ServeConn在该连接使用JSON编解码格式。
		*/
		jsonrpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
