package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	product "go-micro-shopping/product/proto/product"
)

type Product struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Product) Call(ctx context.Context, req *product.Request, rsp *product.Response) error {
	log.Log("Received Product.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Product) Stream(ctx context.Context, req *product.StreamingRequest, stream product.Product_StreamStream) error {
	log.Logf("Received Product.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&product.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Product) PingPong(ctx context.Context, stream product.Product_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&product.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
