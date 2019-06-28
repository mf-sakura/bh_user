package main

import (
	upb "github.com/mf-sakura/bh_user/app/proto"

	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := upb.NewUserServiceClient(conn)
	_, err = c.IncrUserCounter(context.Background(), &upb.IncrUserCounterMessage{
		UserId: 1,
	})
	if err != nil {
		fmt.Printf("error:%v", err)
	}
}
