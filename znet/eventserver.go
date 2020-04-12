package znet

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/xaque208/znet/internal/events"
	"github.com/xaque208/znet/rpc"
	pb "github.com/xaque208/znet/rpc"
)

type eventServer struct {
	eventNames []string
	ch         chan events.Event
	remoteChan chan *rpc.Event
}

func (e *eventServer) ValidEventName(name string) bool {
	for _, n := range e.eventNames {
		if n == name {
			return true
		}
	}

	return false
}

// RegisterEvents is used to update the e.eventNames list.
func (e *eventServer) RegisterEvents(nameSet []string) {
	log.Debugf("eventServer registering %d events: %+v", len(nameSet), nameSet)

	if len(e.eventNames) == 0 {
		e.eventNames = make([]string, 1)
	}

	e.eventNames = append(e.eventNames, nameSet...)
}

// NoticeEvent is the call when an event should be fired.
func (e *eventServer) NoticeEvent(ctx context.Context, request *pb.Event) (*pb.EventResponse, error) {
	response := &pb.EventResponse{}
	e.remoteChan <- request

	if e.ValidEventName(request.Name) {
		ev := events.Event{
			Name:    request.Name,
			Payload: request.Payload,
		}

		e.ch <- ev
	} else {
		response.Errors = true
		response.Message = fmt.Sprintf("unknown RPC event name: %s", request.Name)
		log.Tracef("payload: %s", request.Payload)
		log.Tracef("known events: %+v", e.eventNames)
	}

	return response, nil
}

// SubscribeEvents is used to allow a caller to block while streaming events
// from the event server that match the given event names.
func (e *eventServer) SubscribeEvents(subs *pb.EventSub, stream pb.Events_SubscribeEventsServer) error {
	for ev := range e.remoteChan {

		log.Infof("received remote event %+v", ev)
		match := func() bool {
			for _, n := range subs.Name {
				if n == ev.Name {
					return true
				}
			}

			return false
		}()

		if match {
			log.Debugf("sending remote event: %s", ev.Name)
			if err := stream.Send(ev); err != nil {
				return err
			}
		}
	}

	return nil
}
