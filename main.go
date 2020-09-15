package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2"
	proto "github.com/micro/helloworld/proto"
)

func main() {
	// create and initialise a new service
	service := micro.NewService()
	service.Init()

	// create the proto client for helloworld
	client := proto.NewHelloworldService("go.micro.service.helloworld", service.Client())

	// call an endpoint on the service
	rsp, err := client.Call(context.Background(), &proto.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println("Error calling helloworld: ", err)
		return
	}

	// print the response
	fmt.Println("Response: ", rsp.Msg)

	// let's delay the process for exiting for reasons you'll see below
	time.Sleep(time.Second * 5)
}
