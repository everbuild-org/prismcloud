package server

import (
	"prismcloud.dev/apiserver/services"
	pb "prismcloud.dev/protobufs"
)

type Server struct {
	pb.UnimplementedPrismcloudApiserverServer
	Api services.ServiceApi
}
