package file

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type Srv struct {
}

func (s Srv) GetDataByIndex(ctx context.Context, req *FileReq) (*FileMessage, error) {
	f, _ := os.Open(GetPathFromArgs())
	buf, err := SliceFile(f, 1024, int(req.Index))
	return &FileMessage{
		Index: req.Index,
		Buf:   buf,
	}, err
}

func GrpcServer() {
	listener, err := net.Listen("tcp", ":65010")
	if err != nil {
		fmt.Println(err)
	}
	rpcServer := grpc.NewServer()
	RegisterFileSrvServer(rpcServer, &Srv{})
	reflection.Register(rpcServer)
	if err = rpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
