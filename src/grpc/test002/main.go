package main

import (
	"context"
	"fmt"
	"github.com/sahara-gopher/tour/src/grpc/test002/proto/greet"
	"google.golang.org/grpc"
)

func main() {
	var port = "8999"
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	defer conn.Close()
	client := greet.NewHelloServiceClient(conn)
	rs, err := client.Say(context.Background(), &greet.SayRequest{
		Username: "luoyu",
		Greeting: "love you",
	})
	if err != nil {
		fmt.Printf("response error: %s", err)
	}
	fmt.Println(rs)
}
