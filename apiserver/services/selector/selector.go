package selector

import "strings"

type ServiceType string

const (
	Pod       ServiceType = "pod"
	LBIngress ServiceType = "lbingress"
)

func NamespaceName(namespace string, name string) string {
	return "prismcloud.dev/namespace=" + namespace + ",prismcloud.dev/name=" + name
}

func Managed() string {
	return "prismcloud.dev/managed=true"
}

func Type(t ServiceType) string {
	return "prismcloud.dev/type=" + string(t)
}

func All(values ...string) string {
	return strings.Join(values, ",")
}
