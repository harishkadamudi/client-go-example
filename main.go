package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/hari/.kube/config", "")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		_ = fmt.Errorf(" no such directory %s ", *kubeconfig)
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, pod := range pods.Items {
		fmt.Printf("Pod Name %s", pod.Name)
		fmt.Printf("Pod Status %s", pod.Status)
	}

	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, deploy := range deployments.Items {
		fmt.Printf("Deployment %s", deploy.Name)
	}

}
