package server

import (
	"context"
	pb "prismcloud.dev/protobufs"
)

func (s *Server) CreatePod(_ context.Context, in *pb.PodCreateRequest) (*pb.Void, error) {
	return &pb.Void{}, s.Api.CreatePod(in.Namespace, in.Name, in.Container)
}

func (s *Server) DeletePod(_ context.Context, in *pb.PodDeleteRequest) (*pb.Void, error) {
	return &pb.Void{}, s.Api.DeletePod(in.Namespace, in.Name)
}
