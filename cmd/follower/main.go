// CLI for the TimeMon followers to report their server clock to the coordinator.
package main

import (
	"context"
	"flag"
	"time"

	"github.com/arush15june/TimeMon/pkg/clock"
	"github.com/arush15june/TimeMon/pkg/timeservice"
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	msgID   uint64
	latency uint64

	coordinatorAddr string
	label           string
)

func main() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})

	flag.StringVar(&coordinatorAddr, "coordinator", "127.0.0.1:3530", "Coordinator to connect to")
	flag.StringVar(&label, "label", "follower", "Follower label")

	flag.Parse()

	conn, err := grpc.Dial(coordinatorAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := timeservice.NewTimeServiceClient(conn)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		currClock := clock.GetWallClock()
		tsProto, _ := ptypes.TimestampProto(currClock)

		log.WithFields(log.Fields{
			"msgID":    msgID,
			"type":     "request",
			"currTime": currClock,
			"label":    label,
		}).Info()

		start := time.Now()
		r, err := c.SendClock(ctx, &timeservice.ClockRequest{CurrTime: tsProto, MsgId: msgID, Latency: latency})
		latency = uint64(time.Now().Sub(start).Nanoseconds())

		if err != nil {
			log.Infof("SendClock timeout: %v", err)
		}
		log.WithFields(log.Fields{
			"latency":     latency,
			"msgID_recvd": r.GetMsgId(),
			"response":    r.GetResp(),
			"type":        "response",
			"label":       label,
		}).Info()

		msgID++
		<-time.After(50 * time.Millisecond)
	}

}
