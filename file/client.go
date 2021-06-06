package file

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

// main 方法实现对 gRPC 接口的请求
func GetDataFromServer(address string, index int64) []byte {
	conn, err := grpc.Dial(address+":65010", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Can't connect: " + address)
	}
	defer conn.Close()
	client := NewFileSrvClient(conn)
	resp, err := client.GetDataByIndex(context.Background(), &FileReq{Index: index})
	if err != nil {
		log.Fatalln("Do Format error:" + err.Error())
	}
	return resp.Buf
}
