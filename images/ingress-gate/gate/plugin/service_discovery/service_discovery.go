package service_discovery

import (
	"context"
	"github.com/go-logr/logr"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

var Plugin = proxy.Plugin{
	Name: "service-discovery",
	Init: func(ctx context.Context, proxy *proxy.Proxy) error {
		var logger = logr.FromContextOrDiscard(ctx)
		logger.Info("service-discovery plugin initialized")
		return nil
	},
}
