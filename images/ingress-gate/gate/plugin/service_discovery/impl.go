package service_discovery

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"strings"
)

type DiscoveredService struct {
	Name string
	Pod  string
	Host string
	Port int
}

type ServiceDiscoveryService interface {
	discovery(ctx context.Context) ([]DiscoveredService, error)
}

type NoopServiceDiscoveryService struct{}

func (s *NoopServiceDiscoveryService) discovery(_ context.Context) ([]DiscoveredService, error) {
	return make([]DiscoveredService, 0), nil
}

type KubernetesServiceDiscoveryService struct {
	kubeClient     *kubernetes.Clientset
	namespace      string
	prismNamespace string
}

func (s *KubernetesServiceDiscoveryService) discovery(ctx context.Context) ([]DiscoveredService, error) {
	pods, err := s.kubeClient.CoreV1().Pods(s.namespace).List(ctx, metav1.ListOptions{
		LabelSelector: "prismcloud.dev/namespace=" + s.prismNamespace + ",prismcloud.dev/managed=true,prismcloud.dev/autodiscovery=true",
	})
	if err != nil {
		return nil, err
	}

	var services []DiscoveredService
	for _, pod := range pods.Items {
		labels := pod.GetLabels()
		if labels == nil {
			continue
		}

		// get the prismcloud.dev/name label
		name, ok := labels["prismcloud.dev/name"]
		if !ok {
			continue
		}

		for _, container := range pod.Spec.Containers {
			for _, port := range container.Ports {
				services = append(services, DiscoveredService{
					Name: name,
					Pod:  pod.Name,
					Host: pod.Status.PodIP,
					Port: int(port.ContainerPort),
				})
			}
		}
	}

	return services, nil
}

func parsePrismNamespace() (string, error) {
	// env var __PRISM_NAMESPACE is set by the api server on pod creation
	namespace := os.Getenv("__PRISM_NAMESPACE")

	if namespace == "" {
		return "", fmt.Errorf("no prism namespace found")
	}

	return namespace, nil
}

func newInClusterServiceDiscoveryService() *KubernetesServiceDiscoveryService {
	namespaceFile := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	namespaceBytes, err := os.ReadFile(namespaceFile)
	if err != nil {
		return nil
	}

	namespace := strings.TrimSpace(string(namespaceBytes))

	prismNamespace, err := parsePrismNamespace()

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil
	}

	return &KubernetesServiceDiscoveryService{
		kubeClient:     kubeClient,
		namespace:      namespace,
		prismNamespace: prismNamespace,
	}
}

func NewServiceDiscoveryService(ctx context.Context) ServiceDiscoveryService {
	// check if we are running in a kubernetes cluster
	if _, err := os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		return newInClusterServiceDiscoveryService()
	}

	logger := logr.FromContextOrDiscard(ctx)
	logger.Info("no kubernetes cluster detected, using noop service discovery service")
	return &NoopServiceDiscoveryService{}
}

func (d *DiscoveredService) Network() string {
	return "tcp"
}

func (d *DiscoveredService) String() string {
	return fmt.Sprintf("%s:%d", d.Host, d.Port)
}
