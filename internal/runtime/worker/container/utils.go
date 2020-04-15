package container

import (
	"bytes"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"log"
	"sync/atomic"
)

func Connect(h string) Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	var (
		n1, n2, n3, n4, n5, n6 int64
	)

	return Client{
		h,
		cli,
		stats{
			&n1,
			&n2,
			&n3,
			&n4,
			&n5,
			&n6}}
}

func ReadLogs(logs io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(logs)
	logs.Close()
	return buf.String()
}

func (c *Client) containerStatus(id string) string {
	container, err := c.ContainerInspect(context.Background(), id)

	if err != nil {
		return ""
	}

	return container.State.Status
}

func (c *Client) containerState(id string) *types.ContainerState {
	container, err := c.ContainerInspect(context.Background(), id)

	if err != nil {
		return &types.ContainerState{}
	}

	return container.State
}

func (c *Client) GetMemoryTotal() int64 {
	return atomic.AddInt64(c.memoryTotal, 0)
}

func (c *Client) GetMemoryReserved() int64 {
	return atomic.AddInt64(c.memoryReserved, 0)
}

func (c *Client) GetNanoCPUTotal() int64 {
	return atomic.AddInt64(c.nanoCPUTotal, 0)
}

func (c *Client) GetNanoCPUReserved() int64 {
	return atomic.AddInt64(c.nanoCPUReserved, 0)
}

func (c *Client) GetGPUTotal() int64 {
	return atomic.AddInt64(c.gpuTotal, 0)
}

func (c *Client) GetGPUReserved() int64 {
	return atomic.AddInt64(c.gpuReserved, 0)
}

func (c *Client) SetMemoryTotal(v int64) {
	atomic.StoreInt64(c.memoryTotal, v)
}

func (c *Client) SetNanoCPUTotal(v int64) {
	atomic.StoreInt64(c.nanoCPUTotal, v)
}

func (c *Client) SetGPUTotal(v int64) {
	atomic.StoreInt64(c.gpuTotal, v)
}
