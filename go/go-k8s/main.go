package main

import (
    "fmt"
    "log"
    "path/filepath"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
)

func main() {
    kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        log.Fatalf("Error building kubeconfig: %s", err.Error())
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalf("Error creating Kubernetes client: %s", err.Error())
    }

    version, err := clientset.Discovery().ServerVersion()
    if err != nil {
        log.Fatalf("Error getting server version: %s", err.Error())
    }

    fmt.Printf("Kubernetes server version: %s\n", version.String())
}

