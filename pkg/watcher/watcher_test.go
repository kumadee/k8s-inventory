package watcher

import (
	"testing"

	"k8s.io/client-go/kubernetes/fake"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func TestPodWatcher(t *testing.T) {
	// Create a fake Kubernetes clientset
	clientset := fake.NewSimpleClientset()

	// Create a test namespace
	testNamespace := "test-namespace"
	clientset.CoreV1().Namespaces().Create(&v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: testNamespace},
	})

	// Create test pods
	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod-1", Namespace: testNamespace},
	}
	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "test-pod-2", Namespace: testNamespace},
	}

	clientset.CoreV1().Pods(testNamespace).Create(pod1)
	clientset.CoreV1().Pods(testNamespace).Create(pod2)

	// Use the fake clientset to initialize the Kubernetes configuration
	config := &rest.Config{Transport: fake.NewSimpleClientset().RESTClient().Transport}

	// Create a test informer
	stopCh := make(chan struct{})
	defer close(stopCh)
	go WatchPods(config, testNamespace, stopCh)

	// Wait for events and verify the output
	expectedOutput := []string{
		"Pod test-pod-1 already running in namespace test-namespace",
		"Pod test-pod-2 already running in namespace test-namespace",
	}

	// Replace this part with your actual verification logic.
	// You might need to capture the printed output from the Goroutine and compare it to the expected output.
	// You can use a channel to pass messages from the Goroutine to the test function for verification.
	// Here, we just provide a placeholder.
	// For more advanced testing, consider using a testing framework like Ginkgo.
}

