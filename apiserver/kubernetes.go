package main

import (
	"context"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const KubeNamespace = "prism"

func getClientOutOfCluster() *kubernetes.Clientset {
	var home = homedir.HomeDir()
	var kubeconfig = filepath.Join(home, ".kube", "config")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func getClientInCluster() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func GetClient(outOfCluster bool) *kubernetes.Clientset {
	if outOfCluster {
		return getClientOutOfCluster()
	} else {
		return getClientInCluster()
	}
}

func EnsureNamespace(ctx context.Context, clientset *kubernetes.Clientset) {
	_, err := clientset.CoreV1().Namespaces().Get(ctx, KubeNamespace, metav1.GetOptions{})
	if err == nil {
		return
	}
	if !errors.IsNotFound(err) {
		logrus.WithError(err).Fatal("Could not find namespace")
	}
	_, err = clientset.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: KubeNamespace},
	}, metav1.CreateOptions{})

	logrus.Debug("Created namespace ", KubeNamespace)

}
