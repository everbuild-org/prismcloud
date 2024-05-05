package services

import (
	"context"
	"gorm.io/gorm"

	"k8s.io/client-go/kubernetes"
)

type ServiceApi struct {
	Context   context.Context
	Clientset *kubernetes.Clientset
	Database  *gorm.DB
}
