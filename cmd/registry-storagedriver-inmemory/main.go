package main

import (
	"github.com/BrianBland/docker-registry/storagedriver/inmemory"
	"github.com/BrianBland/docker-registry/storagedriver/ipc"
)

// An out-of-process inmemory driver, intended to be run by ipc.NewDriverClient
// This exists primarily for example and testing purposes
func main() {
	ipc.StorageDriverServer(inmemory.New())
}
