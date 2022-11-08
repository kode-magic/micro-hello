package main

import (
	"hello/handler"
	pb "hello/proto"

	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("hello"),
	)

	// Register handler
	err := pb.RegisterHelloHandler(srv.Server(), handler.New())
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Error(err)
	}
}
