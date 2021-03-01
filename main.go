package main

import (
	"go-category/handler"
	pb "go-category/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("go-category"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterGoCategoryHandler(srv.Server(), new(handler.GoCategory))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
