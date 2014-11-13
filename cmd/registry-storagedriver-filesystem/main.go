package main

import (
	"encoding/json"
	"os"

	"github.com/BrianBland/docker-registry/storagedriver/filesystem"
	"github.com/BrianBland/docker-registry/storagedriver/ipc"
)

// An out-of-process filesystem driver, intended to be run by ipc.NewDriverClient
func main() {
	parametersBytes := []byte(os.Args[1])
	var parameters map[string]string
	err := json.Unmarshal(parametersBytes, &parameters)
	if err != nil {
		panic(err)
	}

	ipc.StorageDriverServer(filesystem.FromParameters(parameters))
}
