package main

import (
	"fmt"
	"golearning/server"
	"log"
	"net/rpc"
)

func main() {
	client, e := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if e != nil {
		log.Fatal("dialing:", e)
	}
	args := &server.Args{57, 8}
	var reply server.Quotient
	err := client.Call("Arith.Devide", args, &reply)
	if err != nil {
		log.Fatal("call error:", err)
	}
	fmt.Println(reply)
}
