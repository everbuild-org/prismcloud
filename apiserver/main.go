package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"prismcloud.dev/apiserver/database"
	"prismcloud.dev/apiserver/server"
	"prismcloud.dev/apiserver/services"
	pb "prismcloud.dev/protobufs"
)

var (
	verbose      = flag.Bool("verbose", false, "Enable verbose logging")
	port         = flag.Int("port", 18948, "The server port")
	outOfCluster = flag.Bool("ooc", false, "Run outside of a cluster (parse kubeconfig file)")
)

func main() {
	flag.Parse()
	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	services.OutOfCluster = *outOfCluster

	logrus.Debug("Connecting to cluster...")

	ctx := context.Background()

	// panics if not configured correctly
	clientset := GetClient(*outOfCluster)
	databaseClient := database.CreateDatabaseClient(*outOfCluster)

	EnsureNamespace(ctx, clientset)

	api := services.ServiceApi{
		Clientset: clientset,
		Database:  databaseClient,
		Context:   ctx,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatal("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterPrismcloudApiserverServer(s, &server.Server{
		Api: api,
	})
	logrus.Info("server listening at ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logrus.Fatal("failed to serve: ", err)
	}
}
