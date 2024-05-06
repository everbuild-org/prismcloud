package services

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"prismcloud.dev/apiserver/services/selector"
	pb "prismcloud.dev/protobufs"
	"time"
)

const KubeNamespace = "prism"

var OutOfCluster bool = false

func (s *ServiceApi) CreatePod(namespace string, name string, container *pb.ContainerConfiguration, autoDiscovery bool) error {
	result, err := s.Clientset.CoreV1().Pods(KubeNamespace).List(s.Context, metav1.ListOptions{
		LabelSelector: selector.All(
			selector.Managed(),
			selector.NamespaceName(namespace, name),
			selector.Type(selector.Pod),
		),
	})

	if err != nil {
		return err
	}

	if len(result.Items) > 0 {
		return fmt.Errorf("pod %s/%s already exists", namespace, name)
	}

	var env []v1.EnvVar
	for k, v := range container.Env {
		env = append(env, v1.EnvVar{
			Name:  k,
			Value: v,
		})
	}

	parsedRam, err := resource.ParseQuantity(fmt.Sprintf("%d", container.Ram))
	if err != nil {
		return err
	}

	pullPolicy := v1.PullAlways
	if OutOfCluster {
		pullPolicy = v1.PullNever
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%s-%d", namespace, name, time.Now().Unix()),
			Namespace: KubeNamespace,
			Labels: map[string]string{
				"prismcloud.dev/managed":       "true",
				"prismcloud.dev/name":          name,
				"prismcloud.dev/namespace":     namespace,
				"prismcloud.dev/type":          string(selector.Pod),
				"prismcloud.dev/port":          fmt.Sprintf("%d", container.Port),
				"prismcloud.dev/autodiscovery": fmt.Sprintf("%t", autoDiscovery),
			},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            name,
					Image:           container.Image,
					ImagePullPolicy: pullPolicy,
					Ports: []v1.ContainerPort{
						{
							ContainerPort: container.Port,
							Protocol:      v1.ProtocolTCP,
						},
					},
					Resources: v1.ResourceRequirements{
						Limits: v1.ResourceList{
							"memory": parsedRam,
						},
					},
					Env: env,
				},
			},
		},
	}

	_, err = s.Clientset.CoreV1().Pods(KubeNamespace).Create(s.Context, pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceApi) GetPod(namespace string, name string) (*v1.Pod, error) {
	result, err := s.Clientset.CoreV1().Pods(KubeNamespace).List(s.Context, metav1.ListOptions{
		LabelSelector: selector.All(
			selector.Managed(),
			selector.NamespaceName(namespace, name),
			selector.Type(selector.Pod),
		),
	})
	if err != nil {
		return nil, err
	}
	if len(result.Items) > 0 {
		return &result.Items[0], nil
	}

	return nil, nil
}

func (s *ServiceApi) DeletePod(namespace string, name string) error {
	pod, err := s.GetPod(namespace, name)
	if err != nil {
		return err
	}

	if pod == nil {
		return fmt.Errorf("pod %s/%s does not exist", namespace, name)
	}

	err = s.Clientset.CoreV1().Pods(KubeNamespace).Delete(s.Context, pod.Name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
