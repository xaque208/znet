// Copyright © 2018 Zach Leslie <code@zleslie.info>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"os"

	"github.com/jedib0t/go-pretty/table"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pb "github.com/xaque208/znet/rpc"
	"github.com/xaque208/znet/znet"
	"google.golang.org/grpc"
)

// invCmd represents the inv command
var invCmd = &cobra.Command{
	Use:   "inv",
	Short: "Report on inventory",
	Long:  "",
	Run:   runInv,
}

var rpcServer string

func init() {
	rootCmd.AddCommand(invCmd)

	invCmd.PersistentFlags().StringVarP(&rpcServer, "rpc", "r", ":8800", "Specify RPC server address")
	invCmd.Flags().BoolP("verbose", "v", false, "Raise verbosity")
}

func runInv(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	z, err := znet.NewZnet(cfgFile)
	if err != nil {
		log.Error(err)
	}

	z.Config.RPC.ServerAddress = viper.GetString("rpc.server")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(z.Config.RPC.ServerAddress, opts...)
	if err != nil {
		log.Error(err)
	}
	defer conn.Close()

	client := pb.NewInventoryClient(conn)

	req := &pb.SearchRequest{}

	res, err := client.Search(context.Background(), req)
	if err != nil {
		log.Error(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Platform", "Type", "Description"})

	for _, h := range res.Hosts {
		t.AppendRow([]interface{}{
			h.CommonName,
			h.Platform,
			h.Type,
			h.Description,
		})
	}

	// t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.Render()

}
