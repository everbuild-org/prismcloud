package resources

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"prismcloud.dev/cli/client"
	"prismcloud.dev/protobufs"
	"strings"
)

type LBIngressYamlDeclaration struct {
	Kind     string `yaml:"kind"`
	Name     string `yaml:"name"`
	Selector struct {
		Pod string `yaml:"pod,omitempty"`
	} `yaml:"selector"`
	Port struct {
		ContainerPort int32  `yaml:"containerPort"`
		ServicePort   int32  `yaml:"servicePort"`
		Protocol      string `yaml:"protocol"`
	}
}

func applyLbIngress(contents string, addr string, namespace string) error {
	var resource LBIngressYamlDeclaration
	if err := yaml.Unmarshal([]byte(contents), &resource); err != nil {
		return err
	}

	c, err := client.NewClient(addr)
	if err != nil {
		return err
	}

	selector := &protobufs.ServiceSelector{
		Type: protobufs.ServiceSelectorType_SERVICE_SELECTOR_TYPE_NONE,
		Name: "",
	}

	if resource.Selector.Pod != "" {
		selector.Type = protobufs.ServiceSelectorType_SERVICE_SELECTOR_TYPE_POD
		selector.Name = resource.Selector.Pod
	}

	if selector.Type == protobufs.ServiceSelectorType_SERVICE_SELECTOR_TYPE_NONE {
		return fmt.Errorf("no valid selector type specified")
	}

	portProtocol, ok := protobufs.IngressProtocol_value[strings.ToUpper(resource.Port.Protocol)]
	if !ok {
		return fmt.Errorf("no valid port protocol specified")
	}

	_, err = c.Api.CreateLBIngress(c.Ctx, &protobufs.LBIngressCreateRequest{
		Name:      resource.Name,
		Namespace: namespace,
		Selector:  selector,
		Port: &protobufs.LBIngressPortConfiguration{
			ContainerPort: resource.Port.ContainerPort,
			ServicePort:   resource.Port.ServicePort,
			Protocol:      protobufs.IngressProtocol(portProtocol),
		},
	})

	if err != nil {
		return err
	}

	fmt.Printf("create lbingress/%s/%s\n", namespace, resource.Name)
	return nil
}

func deleteLbIngress(contents string, addr string, namespace string) error {
	var resource LBIngressYamlDeclaration
	if err := yaml.Unmarshal([]byte(contents), &resource); err != nil {
		return err
	}

	c, err := client.NewClient(addr)
	if err != nil {
		return err
	}

	_, err = c.Api.DeleteLBIngress(c.Ctx, &protobufs.LBIngressDeleteRequest{
		Name:      resource.Name,
		Namespace: namespace,
	})

	if err != nil {
		return err
	}

	fmt.Printf("delete lbingress/%s\n", namespace)

	return nil
}
