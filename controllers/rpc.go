package controllers

import (
	"beego_action/helpers"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcController struct {
	BaseController
}

type Sum struct {
	X, Y int
}

type Args struct {
	A, B int
}

func (this *Sum) Add(args *Args, reply *int) error {
	res := args.A + args.B
	reply = &res
	return nil
}

func (this *RpcController) Prepare() {
	fmt.Println("rpc_controller prepare")
	this.PrepareViewData()

	//this.server_json()
	//this.server_http()
	//this.server_tcp()
	fmt.Println("server start")
}

func (this *RpcController) Client_json() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("rpc_json")
	fmt.Printf("Sum: %d*%d=%d\n", args.A, args.B, reply)

}

func (this *RpcController) server_json() {
	arith := new(Sum)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	helpers.CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	helpers.CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func (this *RpcController) Client_http() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1235")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Sum.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("rpc_http")
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

func (this *RpcController) server_http() {
	arith := new(Sum)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1235", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (this *RpcController) Client_tcp() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1236")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Sum.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println("rpc_tcp")
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

func (this *RpcController) server_tcp() {
	arith := new(Sum)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1236")
	helpers.CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	helpers.CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
