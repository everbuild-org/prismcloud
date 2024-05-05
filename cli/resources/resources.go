package resources

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Action int

const (
	Apply Action = iota
	Delete
)

type KindDeclaration struct {
	Kind             string `yaml:"kind"`
	DefaultNamespace string `yaml:"defaultNamespace,omitempty"`
	Name             string `yaml:"name"`
}

var applyFunctions = map[string]interface{}{
	"Namespace": applyNamespace,
	"Pod":       applyPod,
	"LBIngress": applyLbIngress,
}

var deleteFunctions = map[string]interface{}{
	"Namespace": deleteNamespace,
	"Pod":       deletePod,
	"LBIngress": deleteLbIngress,
}

func ParseAndActOnResourceFile(contents string, addr string, action Action, namespace string) error {
	var kind KindDeclaration
	if err := yaml.Unmarshal([]byte(contents), &kind); err != nil {
		return err
	}

	if namespace == "" && kind.Kind != "Namespace" {
		if kind.DefaultNamespace == "" {
			return fmt.Errorf("no default namespace for kind %s", kind.Kind)
		}

		namespace = kind.DefaultNamespace
	}

	if len(kind.Name) > 22 {
		return fmt.Errorf("invalid resource name length %d", len(kind.Name))
	}

	if action == Apply {
		if f, ok := applyFunctions[kind.Kind]; ok {
			return f.(func(string, string, string) error)(contents, addr, namespace)
		}

		return fmt.Errorf("unknown resource action: %s", kind.Kind)
	} else if action == Delete {
		if f, ok := deleteFunctions[kind.Kind]; ok {
			return f.(func(string, string, string) error)(contents, addr, namespace)
		}

		return fmt.Errorf("unknown resource action: %s", kind.Kind)
	} else {
		return fmt.Errorf("unknown resource action: %s", kind.Kind)
	}
}
