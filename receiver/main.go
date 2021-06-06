package main
import (
	"fmt"
	"syscall"
	"multicast-file/net"
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
	net.UDPMulticast(socketMC)
	net.RecvMsg(socketMC)

}
