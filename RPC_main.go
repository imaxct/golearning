package main

import (
	"net/rpc"
	"net"
	"log"
	"net/http"
	"golearning/server"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
