package main

import "github.com/kumadee/k8s-inventory/pkg/watcher"
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
  // Initialize a context that can be used for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a channel to capture interrupt signals (e.g., SIGINT, SIGTERM)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// Block until a signal is received
		<-signalCh
		fmt.Println("Received shutdown signal. Cleaning up...")
		cancel() // Trigger graceful shutdown
	}()

  client := watcher.GetClient()
  namespace := "default"

  watcher.PodWatcher(namespace, client)

  // Wait for the context to be canceled (shutdown signal received)
	<-ctx.Done()
	fmt.Println("Shutdown complete.")
}
