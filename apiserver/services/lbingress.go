package services

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"prismcloud.dev/apiserver/services/selector"
	sel "prismcloud.dev/apiserver/services/selector"
	"prismcloud.dev/protobufs"
	"strconv"
	"time"
)

func (s *ServiceApi) CreateLBIngress(namespace string, name string, selector *protobufs.ServiceSelector, port *protobufs.LBIngressPortConfiguration) error {
	if port == nil || selector == nil {
		return fmt.Errorf("port or selector is nil")
	}

	result, err := s.Clientset.CoreV1().Services(KubeNamespace).List(s.Context, metav1.ListOptions{
		LabelSelector: sel.All(
			sel.Managed(),
			sel.NamespaceName(namespace, name),
			sel.Type(sel.LBIngress),
		),
	})

	if err != nil {
		return err
	}

	if len(result.Items) > 0 {
		return fmt.Errorf("lbingress %s/%s already exists", namespace, name)
	}

	var typeStr string
	switch selector.Type {
	case protobufs.ServiceSelectorType_SERVICE_SELECTOR_TYPE_NONE:
		return fmt.Errorf("no valid selector type specified")
	case protobufs.ServiceSelectorType_SERVICE_SELECTOR_TYPE_POD:
		typeStr = string(sel.Pod)
	}

	_, err = s.Clientset.CoreV1().Services(KubeNamespace).Create(s.Context, &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%s-%d", namespace, name, time.Now().Unix()),
			Namespace: KubeNamespace,
			Labels: map[string]string{
				"prismcloud.dev/managed":     "true",
				"prismcloud.dev/name":        name,
				"prismcloud.dev/namespace":   namespace,
				"prismcloud.dev/type":        string(sel.LBIngress),
				"prismcloud.dev/port":        fmt.Sprintf("%d", port.ContainerPort),
				"prismcloud.dev/servicePort": fmt.Sprintf("%d", port.ServicePort),
				"prismcloud.dev/protocol":    protobufs.IngressProtocol_name[int32(port.Protocol)],
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"prismcloud.dev/managed":   "true",
				"prismcloud.dev/name":      selector.Name,
				"prismcloud.dev/namespace": namespace,
				"prismcloud.dev/type":      typeStr,
			},
			Ports: []v1.ServicePort{
				{
					Name:       "http",
					Port:       port.ServicePort,
					TargetPort: intstr.FromInt32(port.ContainerPort),
					Protocol:   v1.Protocol(protobufs.IngressProtocol_name[int32(port.Protocol)]),
				},
			},
			Type: v1.ServiceTypeLoadBalancer,
		},
	}, metav1.CreateOptions{})

	return err
}

func (s *ServiceApi) DeleteLBIngress(namespace string, name string) error {
	_, ingress, err := s.GetLBIngress(namespace, name)
	if err != nil {
		return err
	}

	if ingress == nil {
		return fmt.Errorf("lbingress %s/%s does not exist", namespace, name)
	}

	return s.Clientset.CoreV1().Services(KubeNamespace).Delete(s.Context, ingress.Name, metav1.DeleteOptions{})
}

func (s *ServiceApi) GetLBIngress(namespace string, name string) (*protobufs.LBIngress, *v1.Service, error) {
	result, err := s.Clientset.CoreV1().Services(KubeNamespace).List(s.Context, metav1.ListOptions{
		LabelSelector: selector.All(
			selector.Managed(),
			selector.NamespaceName(namespace, name),
			selector.Type(selector.LBIngress),
		),
	})

	if err != nil {
		return nil, nil, err
	}

	if len(result.Items) > 0 {
		item := result.Items[0]
		itemName := item.Labels["prismcloud.dev/name"]
		itemNamespace := item.Labels["prismcloud.dev/namespace"]
		itemContainerPortStr := item.Labels["prismcloud.dev/port"]
		itemServicePortStr := item.Labels["prismcloud.dev/servicePort"]
		itemProtocolStr := item.Labels["prismcloud.dev/protocol"]

		itemContainerPort, err := strconv.Atoi(itemContainerPortStr)
		if err != nil {
			return nil, nil, err
		}

		itemServicePort, err := strconv.Atoi(itemServicePortStr)
		if err != nil {
			return nil, nil, err
		}

		itemProtocol := protobufs.IngressProtocol(protobufs.IngressProtocol_value[itemProtocolStr])

		return &protobufs.LBIngress{
			Name:          itemName,
			Namespace:     itemNamespace,
			ContainerPort: int32(itemContainerPort),
			ServicePort:   int32(itemServicePort),
			Protocol:      itemProtocol,
		}, &item, nil
	}

	return nil, nil, nil
}
