package main

import (
	"golearning/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
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
