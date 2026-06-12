package main

import (
	"context"
	"fmt"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// Load kubeconfig
	kubeconfig := filepath.Join(
		"/Users/rohan/.kube",
		"config",
	)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Get pods
	pods, err := clientset.CoreV1().
		Pods("test").
		List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Pods in default namespace:")

	for _, pod := range pods.Items {
		printPod(pod)
	}
}

func printPod(pod v1.Pod) {
	fmt.Printf("- %s\n", pod.Name)
}
