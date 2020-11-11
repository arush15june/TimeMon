// CLI for the TimeMon followers to report their server clock to the coordinator.
package main

import (
	"context"
	"time"

	"github.com/arush15june/TimeMon/pkg/clock"
	"github.com/arush15june/TimeMon/pkg/timeservice"
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	coordinatorAddr = "localhost:3530"
)

var (
	msgId uint64 = 0
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.Info("Starting follower")

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

		log.Infof("Sending %s", currClock)

		r, err := c.SendClock(ctx, &timeservice.ClockRequest{CurrTime: tsProto, MsgId: msgId})

		if err != nil {
			log.Fatalf("SendClock timeout: %v", err)
		}
		log.Printf("Response: %s %d", r.GetResp(), r.GetMsgId())

		msgId += 1
		<-time.After(250 * time.Millisecond)
	}

}
