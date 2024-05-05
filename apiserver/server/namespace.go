package server

import (
	"context"
	"fmt"
	"prismcloud.dev/apiserver/database/model"

	pb "prismcloud.dev/protobufs"
)

func (s *Server) CreateNamespace(_ context.Context, in *pb.NamespaceCreateRequest) (*pb.Namespace, error) {
	namespace := &model.Namespace{
		Name:     in.Name,
		RamLimit: in.RamLimit,
	}

	if !s.Api.HasNamespace(namespace.Name) {
		if err := s.Api.CreateNamespace(namespace.Name, namespace.RamLimit); err != nil {
			return nil, err
		}
	} else {
		if err := s.Api.PatchNamespace(namespace.Name, namespace.RamLimit); err != nil {
			return nil, err
		}
	}

	return &pb.Namespace{
		Name:     namespace.Name,
		RamLimit: namespace.RamLimit,
	}, nil
}

func (s *Server) GetNamespaces(_ context.Context, _ *pb.Void) (*pb.Namespaces, error) {
	namespaces, err := s.Api.ListNamespaces()
	if err != nil {
		return nil, err
	}

	list := &pb.Namespaces{}
	for _, namespace := range namespaces {
		list.Namespaces = append(list.Namespaces, &pb.Namespace{
			Name:     namespace.Name,
			RamLimit: namespace.RamLimit,
		})
	}

	return list, nil
}

func (s *Server) DeleteNamespace(_ context.Context, in *pb.NamespaceDeleteRequest) (*pb.Void, error) {
	if !s.Api.HasNamespace(in.Name) {
		return nil, fmt.Errorf("namespace '%v' does not exist", in.Name)
	}

	if err := s.Api.DeleteNamespace(in.Name); err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}
