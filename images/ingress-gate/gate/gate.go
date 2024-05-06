package gate

import (
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"prismcloud.dev/gate-ingress/gate/plugin/service_discovery"
)

func Execute() {
	proxy.Plugins = append(proxy.Plugins,
		service_discovery.Plugin,
	)

	gate.Execute()
}
