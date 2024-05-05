package resources

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"prismcloud.dev/cli/client"
	"prismcloud.dev/protobufs"
)

type PodYamlDeclaration struct {
	Kind      string `yaml:"kind"`
	Name      string `yaml:"name"`
	Container struct {
		Image string            `yaml:"image"`
		Port  int32             `yaml:"port"`
		Ram   string            `yaml:"ram"`
		Env   map[string]string `yaml:"env,omitempty"`
	} `yaml:"container"`
}

func applyPod(contents string, addr string, namespace string) error {
	var pod PodYamlDeclaration
	if err := yaml.Unmarshal([]byte(contents), &pod); err != nil {
		return err
	}

	ram, err := client.ParseRam(pod.Container.Ram)
	if err != nil {
		return err
	}

	if pod.Container.Env == nil {
		pod.Container.Env = make(map[string]string)
	}

	c, err := client.NewClient(addr)
	if err != nil {
		return err
	}

	_, err = c.Api.CreatePod(c.Ctx, &protobufs.PodCreateRequest{
		Name:      pod.Name,
		Namespace: namespace,
		Container: &protobufs.ContainerConfiguration{
			Image: pod.Container.Image,
			Port:  pod.Container.Port,
			Ram:   ram,
			Env:   pod.Container.Env,
		},
	})

	if err != nil {
		return err
	}

	fmt.Printf("create pod/%s/%s\n", namespace, pod.Name)
	return nil
}

func deletePod(contents string, addr string, namespace string) error {
	var pod PodYamlDeclaration
	if err := yaml.Unmarshal([]byte(contents), &pod); err != nil {
		return err
	}

	c, err := client.NewClient(addr)
	if err != nil {
		return err
	}

	_, err = c.Api.DeletePod(c.Ctx, &protobufs.PodDeleteRequest{
		Name:      pod.Name,
		Namespace: namespace,
	})

	if err != nil {
		return err
	}

	fmt.Printf("delete pod/%s\n", namespace)

	return nil
}
