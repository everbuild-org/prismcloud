package service_discovery

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/robinbraemer/event"
	gateProxy "go.minekube.com/gate/pkg/edition/java/proxy"
	"k8s.io/utils/set"
	"math/rand"
	"strings"
	"time"
)

func discoverLoop(ctx context.Context, discoveryService ServiceDiscoveryService, proxy *gateProxy.Proxy, logger logr.Logger) {
	for {
		services, err := discoveryService.discovery(ctx)
		if err != nil {
			logger.Error(err, "discovery failed")
			time.Sleep(5 * time.Second)
			continue
		}

		currentServices := make(set.Set[string])
		for _, service := range proxy.Servers() {
			currentServices.Insert(service.ServerInfo().Name())
		}

		for _, service := range services {
			if server := proxy.Server(service.Name); server != nil {
				if server.ServerInfo().Addr().String() != service.String() {
					logger.Info("updating service", "name", service.Name, "pod", service.Pod, "host", service.Host, "port", service.Port)
					proxy.Unregister(server.ServerInfo())
					serverInfo := gateProxy.NewServerInfo(service.Name, &service)
					_, err := proxy.Register(serverInfo)
					if err != nil {
						logger.Error(err, "failed to re-register service", "name", service.Name, "pod", service.Pod, "host", service.Host, "port", service.Port)
					}
				}
				currentServices.Delete(service.Name)
				continue
			}

			serverInfo := gateProxy.NewServerInfo(service.Name, &service)
			_, err := proxy.Register(serverInfo)
			if err != nil {
				logger.Error(err, "failed to register service", "name", service.Name, "pod", service.Pod, "host", service.Host, "port", service.Port)
				continue
			}

			currentServices.Delete(service.Name)
		}

		for service := range currentServices {
			logger.Info("unregistering service", "name", service)
			server := proxy.Server(service)
			if server == nil {
				continue // server was already unregistered
			}

			proxy.Unregister(server.ServerInfo())
		}

		// sleep for a bit
		time.Sleep(time.Second * 3)
	}
}

func subscribeToEvents(proxy *gateProxy.Proxy, logger logr.Logger) {
	event.Subscribe(proxy.Event(), 0, func(e *gateProxy.PlayerChooseInitialServerEvent) {
		// Try to connect to a service named lobby (if multiple, pick a random one)
		var lobbyServers []gateProxy.RegisteredServer
		for _, server := range proxy.Servers() {
			name := server.ServerInfo().Name()
			if strings.HasPrefix(strings.ToLower(name), "lobby") {
				lobbyServers = append(lobbyServers, server)
			}
		}

		if len(lobbyServers) == 0 {
			if len(proxy.Servers()) == 0 {
				return
			}

			// Pick first server
			e.SetInitialServer(proxy.Servers()[0])
		} else {
			length := len(lobbyServers)
			if length == 1 {
				e.SetInitialServer(lobbyServers[0])
			} else {
				index := rand.Intn(length)
				e.SetInitialServer(lobbyServers[index])
			}
		}
	})
}

var Plugin = gateProxy.Plugin{
	Name: "service-discovery",
	Init: func(ctx context.Context, proxy *gateProxy.Proxy) error {
		var logger = logr.FromContextOrDiscard(ctx)
		var discoveryService = NewServiceDiscoveryService(ctx)
		go discoverLoop(ctx, discoveryService, proxy, logger)
		subscribeToEvents(proxy, logger)

		logger.Info("service-discovery plugin initialized")
		return nil
	},
}
