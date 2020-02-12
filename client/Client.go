package main

import (
	"bufio"
	"context"
	"fmt"
	chat "github.com/elessarwin/Simple-Chat-Server/service/proto"
	"google.golang.org/grpc"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Must have a url to connect to as the first argument, and a username as the second argument")
		return
	}

	ctx := context.Background()

	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := chat.NewChatServerClient(conn)
	stream, err := c.Chat(ctx)
	if err != nil {
		panic(err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			} else if err != nil {
				panic(err)
			}
			fmt.Println(msg.UserId + ": " + msg.Message)
		}
	}()

	fmt.Println("Connection established, type \"quit\" or use ctrl+c to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "quit" {
			err := stream.CloseSend()
			if err != nil {
				panic(err)
			}
			break
		}

		err := stream.Send(&chat.ChatMessage{
			UserId:  os.Args[2],
			Message: msg,
		})
		if err != nil {
			panic(err)
		}
	}

	<-waitc
}
