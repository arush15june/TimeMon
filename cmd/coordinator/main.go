// CLI for the TimeMon followers to report their server clock to the coordinator.
package main

import (
	"context"
	"flag"
	"net"
	"sync/atomic"
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

	followerLabel := in.GetLabel()

	if followerLabel == "follower-1" {
		atomic.StoreInt64(&followerOneTimeStamp, callerTime.UnixNano())
		atomic.StoreUint64(&followerOneLatency, in.GetLatency())
	} else if followerLabel == "follower-2" {
		atomic.StoreInt64(&followerTwoTimeStamp, callerTime.UnixNano())
		atomic.StoreUint64(&followerTwoLatency, in.GetLatency())
	}

	log.WithFields(log.Fields{
		"caller":        callerAddr,
		"callerTs":      callerTime.UnixNano(),
		"msgId":         in.GetMsgId(),
		"ts":            time.Now().UnixNano(),
		"f1_ts":         followerOneTimeStamp,
		"f2_ts":         followerTwoTimeStamp,
		"f1_latency":    followerOneLatency,
		"f2_latency":    followerTwoLatency,
		"prev_latency":  in.GetLatency(),
		"followerLabel": followerLabel,
		"label":         label,
	}).Info()

	return &timeservice.ClockReply{Resp: timeservice.Response_SUCCESS, MsgId: in.GetMsgId()}, nil
}

var (
	listenAddr string
	label      string

	followerOneTimeStamp int64 = 0
	followerTwoTimeStamp int64 = 0

	followerOneLatency uint64 = 0
	followerTwoLatency uint64 = 0
)

func main() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	flag.StringVar(&listenAddr, "listenAddr", ":3530", "coordinator listen address")
	flag.StringVar(&label, "label", "coordinator", "coordinator label")

	flag.Parse()

	lis, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	timeservice.RegisterTimeServiceServer(s, &CoordinationServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
