package main

import (
    "context"
    "flag"
    "fmt"
    "os"
    "path/filepath"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
    // kubeconfig path (adjust if necessary)
    kubeconfig := flag.String("kubeconfig", filepath.Join(
        homeDir(), ".kube", "config"), "absolute path to the kubeconfig file")
    flag.Parse()

    // Build config from kubeconfig file
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    // Create clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    // Set namespace and secret name to fetch
    namespace := "sandbox"
    secretName := "platform-postgres-user-env-gt8fc4mgd4"

    // Fetch the secret
    secret, err := clientset.CoreV1().Secrets(namespace).Get(context.Background(), secretName, metav1.GetOptions{})
    if err != nil {
        panic(err.Error())
    }

    // Print secret data keys and values (be careful with sensitive info)
    fmt.Printf("Secret %q in namespace %q:\n", secretName, namespace)
    for k, v := range secret.Data {
        fmt.Printf("  %s: %s\n", k, string(v))
    }
}

func homeDir() string {
    if h := os.Getenv("HOME"); h != "" {
        return h
    }
    return os.Getenv("USERPROFILE") // windows fallback
}

