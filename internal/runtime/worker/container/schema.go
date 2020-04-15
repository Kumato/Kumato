package container

import (
	"github.com/docker/docker/client"
)

const (
	EnvKeyInternalToken     = "INTERNAL_TOKEN"
	EnvKeyControllerAddress = "CONTROLLER_ADDRESS"
	EnvKeyFileID            = "FILE_ID"
	EnvKeyTaskID            = "TASK_ID"
)

type stats struct {
	memoryTotal,
	memoryReserved,
	nanoCPUTotal,
	nanoCPUReserved,
	gpuTotal,
	gpuReserved *int64
}

type Client struct {
	host string
	*client.Client
	stats
}
