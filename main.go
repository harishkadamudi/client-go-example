package main

import (
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/hari/.kube/config", "")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		_ = fmt.Errorf(" no such directory %s ", *kubeconfig)
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Errorf("error %s, getting incluster configuration", err.Error())
			panic(err)
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) {
			fmt.Println(" Add called")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println(" Delete called")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("deleted was called")
		},
	})
	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)
	pod, err := podInformer.Lister().Pods("default").Get("default")
	fmt.Println(pod)

}
