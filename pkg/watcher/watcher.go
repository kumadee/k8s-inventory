package watcher

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/homedir"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PodWatcher(namespace: str, clientset: *DynamicClient) {
	// Fetch the list of pods in the specified namespace
	podList, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range podList.Items {
		fmt.Printf("Pod %s already running in namespace %s\n", pod.Name, pod.Namespace)
	}

	// Set up a shared informer to watch for pod events
	podInformer := cache.NewSharedInformer(
		cache.NewListWatchFromClient(
			clientset.CoreV1().RESTClient(),
			"pods",
			namespace,
			cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {
					pod := obj.(*v1.Pod)
					fmt.Printf("Pod %s created in namespace %s\n", pod.Name, pod.Namespace)
				},
				DeleteFunc: func(obj interface{}) {
					pod := obj.(*v1.Pod)
					fmt.Printf("Pod %s deleted from namespace %s\n", pod.Name, pod.Namespace)
				},
			},
		),
		time.Second*0, // Resync every 0 seconds
	)

  stopCh := make(chan struct{})
	defer close(stopCh)

	go func() {
		// Start the informer and gracefully exit when the context is done
		defer close(stopCh)
		podInformer.Run(stopCh)
	}()
}

