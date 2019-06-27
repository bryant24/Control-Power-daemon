package main

import (
	"log"
	"net"
	"net/rpc"
	"os/exec"
)

var VERSION = "1.0.0.626"

type Pmanager int
type Args struct {
	Op string
}

func main() {
	power := new(Pmanager)
	rpc.Register(power)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	log.Println("成功监听端口:1234")

	listener, err := net.ListenTCP("tcp", tcpAddr)
	log.Println("节点Rpc服务启动成功,版本号:", VERSION)
	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}

}

func (pm *Pmanager) Do(args *Args, reply *bool) (err error) {
	switch args.Op {
	case "restart":
		err = exec.Command("shutdown", `-r`, `-t`, `0`).Run()
		log.Println("restart:", err)
	case "shutdown":
		err = exec.Command("shutdown", `-s`, `-t`, `0`).Run()
		log.Println("shutdown:", err)
	}
	*reply = true
	return nil

}
