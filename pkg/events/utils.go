package events

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/xaque208/znet/rpc"
)

func MakeRPCEvent(t interface{}) *rpc.Event {
	payload, err := json.Marshal(t)
	if err != nil {
		log.Error(err)
	}

	req := &rpc.Event{
		Name:    reflect.TypeOf(t).Name(),
		Payload: payload,
	}

	return req
}

func ProduceEvent(conn *grpc.ClientConn, ev interface{}) error {
	if conn == nil {
		return fmt.Errorf("unable to make use of nil grpc client")
	}

	ec := rpc.NewEventsClient(conn)
	// t := reflect.TypeOf(ev).String()

	req := MakeRPCEvent(ev)

	if req == nil {
		return fmt.Errorf("failed to make event")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// log.Tracef("producing RPC event %+v", req)
	res, err := ec.NoticeEvent(ctx, req)
	if err != nil {
		return err
	}

	if res.Errors {
		return errors.New(res.Message)
	}

	return nil
}
