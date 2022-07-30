package k8s

import (
	"log"
	"os"
	"path/filepath"

	"github.com/JYisus/PodChaosMonkey/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewClientset(cfg *config.Config) *kubernetes.Clientset {
	kubeConfig := newKubernetesConfig(cfg)

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("Fatal error creating kubernetes clientset: %s", err)
	}

	return clientset
}

func newKubernetesConfig(cfg *config.Config) *rest.Config {
	if cfg.IsInsideCluster {
		log.Println("Using Kubernetes in-cluster configuration")

		kubeConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Fatal error retrieving kubernetes configuration: %s", err)
		}

		return kubeConfig
	}

	log.Println("Using Kubernetes local configuration")

	kubeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf("Fatal error creating Kubernetes config: %s", err)
	}

	return kubeConfig
}
