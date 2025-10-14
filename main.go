package main

import (
	"fmt"
	"log"
	"net"

	"github.com/relaunch-cot/service-chat/config"
	"github.com/relaunch-cot/service-chat/resource"
	"github.com/relaunch-cot/service-chat/server/methods"
	"google.golang.org/grpc"
)

func main() {
	resource.Inject()

	lis, err := net.Listen("tcp", ":"+config.PORT)
	fmt.Println("Listening on " + config.PORT)
	if err != nil {
		log.Fatalf("Failed to listen on %v: %v\n", config.PORT, err)
	}

	s := grpc.NewServer()

	methods.RegisterGrpcServices(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
