package main

import (
	"fmt"
	"github.com/elessarwin/Simple-Chat-Server/Server/models"
	chat "github.com/elessarwin/Simple-Chat-Server/service/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	srv := models.NewChatServer()
	chat.RegisterChatServerServer(s, srv)

	fmt.Println("Now serving at port 8080")
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}
}
