package server

import (
	"context"
	pb "prismcloud.dev/protobufs"
)

func (s *Server) CreateLBIngress(_ context.Context, in *pb.LBIngressCreateRequest) (*pb.LBIngress, error) {
	err := s.Api.CreateLBIngress(in.Namespace, in.Name, in.Selector, in.Port)
	if err != nil {
		return nil, err
	}

	return &pb.LBIngress{
		Name:          in.Name,
		Namespace:     in.Namespace,
		ContainerPort: in.Port.ContainerPort,
		ServicePort:   in.Port.ServicePort,
		Protocol:      in.Port.Protocol,
	}, nil
}

func (s *Server) DeleteLBIngress(_ context.Context, in *pb.LBIngressDeleteRequest) (*pb.Void, error) {
	err := s.Api.DeleteLBIngress(in.Namespace, in.Name)

	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (s *Server) GetLBIngress(_ context.Context, in *pb.ServiceFilter) (*pb.LBIngress, error) {
	ingress, _, err := s.Api.GetLBIngress(in.Namespace, in.Name)
	return ingress, err
}
