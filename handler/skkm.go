package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	skkm "skkm/proto/skkm"
)

type Skkm struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Skkm) Call(ctx context.Context, req *skkm.Request, rsp *skkm.Response) error {
	log.Info("Received Skkm.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Skkm) Stream(ctx context.Context, req *skkm.StreamingRequest, stream skkm.Skkm_StreamStream) error {
	log.Infof("Received Skkm.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&skkm.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Skkm) PingPong(ctx context.Context, stream skkm.Skkm_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&skkm.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
