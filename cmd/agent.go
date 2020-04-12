package cmd

import (
	"context"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/xaque208/znet/internal/events"
	pb "github.com/xaque208/znet/rpc"
	"github.com/xaque208/znet/znet"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Run a znet agent",
	Run:   runAgent,
}

func init() {
	rootCmd.AddCommand(agentCmd)
}

func runAgent(cmd *cobra.Command, args []string) {
	formatter := log.TextFormatter{
		FullTimestamp: true,
	}

	log.SetFormatter(&formatter)
	if trace {
		log.SetLevel(log.TraceLevel)
	} else if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	z, err := znet.NewZnet(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	z.Config.RPC.ServerAddress = viper.GetString("rpc.server")

	consumers := []events.Consumer{}

	err = z.EventMachine(consumers)
	if err != nil {
		log.Error(err)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(z.Config.RPC.ServerAddress, opts...)
	if err != nil {
		log.Error(err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Error(err)
		}
	}()

	client := pb.NewEventsClient(conn)

	eventSub := &pb.EventSub{
		Name: []string{
			"NewCommit",
			"NewTag",
		},
	}

	stream, err := client.SubscribeEvents(context.Background(), eventSub)
	if err != nil {
		log.Fatalf("calling %+v.SubscribeEvents(_) = _, %+v", client, err)
	}

	for {
		ev, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Errorf("%v.SubscribeEvents(_) = _, %v", client, err)
			break
		}

		log.Tracef("received event: %+v", ev)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		log.Warnf("caught signal: %s", sig.String())
		done <- true
	}()

	<-done
}