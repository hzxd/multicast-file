package main

import (
	"flag"
	"fmt"
	"multicast-file/file"
	"multicast-file/net"
	"syscall"
	"time"
)

func main() {
	socketMC, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		fmt.Printf("Create socket fd error:%s\n", err.Error())
		return
	}
	defer syscall.Close(socketMC)
	if err = net.SetTTL(socketMC, -1); err != nil {
		fmt.Printf("Set ttl error:%s\n", err.Error())
		return
	}
	path := flag.Args()[len(flag.Args())-1]
	ch := make(chan *file.FileMessage, 10)
	stop := make(chan bool)
	go file.GrpcServer()
	go file.SendFile(path, ch, stop)

	for {
		select {
		case data := <-ch:
			net.SendFileMsg(socketMC, data)
			time.Sleep(time.Millisecond * 100)
			fmt.Println("send ", data.Index)
		case <-stop:
			break
		}
	}

	time.Sleep(time.Second * 5)
}
