// CLI for the TimeMon followers to report their server clock to the coordinator.
package main

import (
	"context"
	"net"
	"time"

	"github.com/arush15june/TimeMon/pkg/timeservice"
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type CoordinationServer struct {
	timeservice.UnimplementedTimeServiceServer
}

func (s *CoordinationServer) SendClock(ctx context.Context, in *timeservice.ClockRequest) (*timeservice.ClockReply, error) {
	p, _ := peer.FromContext(ctx)
	callerAddr := p.Addr.String()

	callerTime, _ := ptypes.Timestamp(in.GetCurrTime())

	log.WithFields(log.Fields{
		"caller":   callerAddr,
		"callerTs": callerTime.UnixNano(),
		"msgId":    in.GetMsgId(),
		"ts":       time.Now().UnixNano(),
	}).Info()

	return &timeservice.ClockReply{Resp: timeservice.Response_SUCCESS, MsgId: in.GetMsgId()}, nil
}

var (
	port string = ":3530"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	timeservice.RegisterTimeServiceServer(s, &CoordinationServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
