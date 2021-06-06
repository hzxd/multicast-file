package net

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"multicast-file/file"
	"os"
	"sync"

	"net"
	"syscall"
)

var (
	//ttl        int
	daemon     bool
	port       = 12389
	multiaddr  = [4]byte{224, 0, 1, 1}
	inaddr_any = [4]byte{0, 0, 0, 0}
	blockMap   = sync.Map{}
)

func SendFileMsg(socketMC int, file *file.FileMessage) error {
	b, err := proto.Marshal(file)
	if err != nil {
		return err
	}
	return SendMsg(socketMC, b)
}

func SendMsg(socketMC int, msg []byte) error {
	if len(msg) == 0 {
		return errors.New("length is zero")
	}
	var (
		raddr = &syscall.SockaddrInet4{Port: port, Addr: multiaddr}
	)

	if err := syscall.Sendto(socketMC, msg, 0, raddr); err != nil {
		return err
	}
	return nil
}

func RecvMsg(socketMC int) {
	var buf = make([]byte, 1024*1024*10)
	var filename string
	var block int64
	var fileMap = make(map[int64]string)
	for {
		n, addr, e := syscall.Recvfrom(socketMC, buf, 0)

		if e != nil {
			break
		}

		fileMsg := file.FileMessage{}
		proto.Unmarshal(buf[:n], &fileMsg)
		go fmt.Printf("Recv:%s %d %d\n", string(fileMsg.Filename), fileMsg.Index, len(fileMap))
		if filename != fileMsg.Filename || block != fileMsg.Block || fileMsg.Type == 1 {
			fileMap = make(map[int64]string)
			filename = fileMsg.Filename
			block = fileMsg.Block
		}
		if fileMsg.Type == 2 {
			raddr, _ := addr.(*syscall.SockaddrInet4)
			fmt.Println(len(fileMap))
			//for k, _ := range fileMap {
			//
			//}
			// 合并文件
			err := file.ConcatFile(filename, fileMap, int(block),
				fmt.Sprintf("%d.%d.%d.%d", raddr.Addr[0], raddr.Addr[1], raddr.Addr[2], raddr.Addr[3]))
			if err != nil {
				fmt.Println(err)
			}
		}
		fname := fmt.Sprintf("%s/%s_%d", os.TempDir(), fileMsg.Filename, fileMsg.Index)
		fileMap[fileMsg.Index] = fname
		go func(buf []byte) {
			fd, _ := os.Create(fname)
			fd.Write(buf)
		}(fileMsg.Buf)

		//if ok {
		//	fmt.Printf("Remote addr:%d.%d.%d.%d:%d\n", raddr.Addr[0], raddr.Addr[1], raddr.Addr[2], raddr.Addr[3], raddr.Port)
		//} else {
		//	fmt.Printf("Remote info:%v\n", addr)
		//}
	}
	ExitMultiCast(socketMC)
}

//加入组播域
func UDPMulticast(socketMC int) error {
	err := syscall.Bind(socketMC, &syscall.SockaddrInet4{Port: port, Addr: inaddr_any})
	if err == nil {
		var mreq = &syscall.IPMreq{Multiaddr: multiaddr, Interface: inaddr_any}
		err = syscall.SetsockoptIPMreq(socketMC, syscall.IPPROTO_IP, syscall.IP_ADD_MEMBERSHIP, mreq)
	}
	return err
}

//退出组播域
func ExitMultiCast(socketMC int) {
	var mreq = &syscall.IPMreq{Multiaddr: multiaddr, Interface: inaddr_any}
	syscall.SetsockoptIPMreq(socketMC, syscall.IPPROTO_IP, syscall.IP_DROP_MEMBERSHIP, mreq)
}

//设置路由的TTL值
func SetTTL(fd, ttl int) error {
	if ttl == -1 {
		ttl = 8
	}
	return syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_MULTICAST_TTL, ttl)
}

//检查是否是有效的组播地址范围
func CheckMultiCast(addr [4]byte) bool {
	if addr[0] > 239 || addr[0] < 224 {
		return false
	}
	if addr[2] == 0 {
		return addr[3] <= 18
	}
	return true
}

func init() {
	multi := flag.String("m", "224.0.1.1", "-m 224.0.1.1 specify multicast address")
	flag.IntVar(&port, "p", 12389, "-p 12389 specify multi address listen port")
	//flag.IntVar(&ttl, "t", 8, "-t 8 specify ttl value")
	flag.BoolVar(&daemon, "d", false, "-d is a recv client")
	flag.Parse()
	ip := net.ParseIP(*multi)
	if ip != nil {
		copy(multiaddr[:], ip[12:16])
		if CheckMultiCast(multiaddr) {
			return
		}
	}
	fmt.Println("Isvalid multi address")
	syscall.Exit(1)
}
