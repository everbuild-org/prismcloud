package server

import (
	"context"

	pb "prismcloud.dev/protobufs"
)

func (s *Server) Version(_ context.Context, _ *pb.Void) (*pb.ApiVersion, error) {
	kubeversion, err := s.Api.Clientset.ServerVersion()
	if err != nil {
		return nil, err
	}

	return &pb.ApiVersion{
		Major:      1,
		Minor:      0,
		Patch:      0,
		Api:        1,
		Kubernetes: kubeversion.String(),
	}, nil
}
