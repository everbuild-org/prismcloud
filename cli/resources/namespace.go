package resources

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"prismcloud.dev/cli/client"
	"prismcloud.dev/protobufs"
)

type NamespaceYamlDeclaration struct {
	Kind     string `yaml:"kind"`
	Name     string `yaml:"name"`
	RamLimit string `yaml:"ramLimit"`
}

func applyNamespace(contents string, addr string, _ string) error {
	var namespace NamespaceYamlDeclaration
	if err := yaml.Unmarshal([]byte(contents), &namespace); err != nil {
		return err
	}

	ramLimit, err := client.ParseRam(namespace.RamLimit)
	if err != nil {
		return err
	}

	c, err := client.NewClient(addr)
	if err != nil {
		return err
	}

	resource, err := c.Api.CreateNamespace(c.Ctx, &protobufs.NamespaceCreateRequest{
		Name:     namespace.Name,
		RamLimit: ramLimit,
	})

	if err != nil {
		return err
	}

	fmt.Printf("create namespace/%s\n", resource.Name)

	return nil
}

func deleteNamespace(contents string, addr string, _ string) error {
	var namespace NamespaceYamlDeclaration
	if err := yaml.Unmarshal([]byte(contents), &namespace); err != nil {
		return err
	}

	c, err := client.NewClient(addr)
	if err != nil {
		return err
	}

	_, err = c.Api.DeleteNamespace(c.Ctx, &protobufs.NamespaceDeleteRequest{
		Name: namespace.Name,
	})

	if err != nil {
		return err
	}

	fmt.Printf("delete namespace/%s\n", namespace.Name)

	return nil
}
