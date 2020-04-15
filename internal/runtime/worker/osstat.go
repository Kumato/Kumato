package worker

import (
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/types"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"time"
)

func (n *Node) getOSStat() (*types.Stats, error) {
	return &types.Stats{
		MemoryTotal:  n.docker.GetMemoryTotal(),
		MemoryUsed:   n.docker.GetMemoryReserved(),
		NanoCpuTotal: n.docker.GetNanoCPUTotal(),
		NanoCpuUsed:  n.docker.GetNanoCPUReserved(),
		GpuTotal:     n.docker.GetGPUTotal(),
		GpuUsed:      n.docker.GetGPUReserved(),
	}, nil
}

func (n *Node) getOSLoad() (*types.Stats, error) {
	memory, err := memory.Get()
	if err != nil {
		return &types.Stats{}, err
	}

	before, err := cpu.Get()
	if err != nil {
		return &types.Stats{}, err
	}
	time.Sleep(600 * time.Millisecond)
	after, err := cpu.Get()
	if err != nil {
		return &types.Stats{}, err
	}
	total := float64(after.Total - before.Total)

	logger.Info("cpu usage:", float64(after.User-before.User+after.System-before.System)/total*100)

	return &types.Stats{
		Hostname:     n.hostname,
		MemoryTotal:  int64(memory.Total),
		MemoryUsed:   int64(memory.Used + memory.Cached),
		NanoCpuTotal: int64(after.CPUCount * 1000000000),
		CpuUsage:     float64(after.User-before.User+after.System-before.System) / total * 100,
	}, nil
}
