package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	gocategory "go-category/proto"
)

type GoCategory struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GoCategory) Call(ctx context.Context, req *gocategory.Request, rsp *gocategory.Response) error {
	log.Info("Received GoCategory.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GoCategory) Stream(ctx context.Context, req *gocategory.StreamingRequest, stream gocategory.GoCategory_StreamStream) error {
	log.Infof("Received GoCategory.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&gocategory.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GoCategory) PingPong(ctx context.Context, stream gocategory.GoCategory_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&gocategory.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
