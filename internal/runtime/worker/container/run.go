package container

import (
	"context"
	docker "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/types"
	"strconv"
	"sync/atomic"
	"time"
)

func (c *Client) Run(t *types.Task, inTkn, cAddr string) (*types.Task, error) {

	atomic.AddInt64(c.nanoCPUReserved, t.GetCpus())
	atomic.AddInt64(c.memoryReserved, t.GetMemory())
	atomic.AddInt64(c.gpuReserved, t.GetGpus())

	status := c.containerStatus(t.GetContainerId())

	if status == "" {
		resp, err := c.ContainerCreate(context.Background(),
			&container.Config{
				Image: t.GetImageRepo() + ":" + t.GetImageTag(),
				Env: []string{
					EnvKeyInternalToken + "=" + inTkn,
					EnvKeyControllerAddress + "=" + cAddr,
					EnvKeyFileID + "=" + t.GetFileUri(),
					EnvKeyTaskID + "=" + strconv.FormatUint(uint64(t.GetId()), 10),
				},
				Tty: true,
				Cmd: []string{"/kumato-init"},
			},
			&container.HostConfig{
				Resources: container.Resources{
					// For future reference, you have to provide the number of CPUs multiplied by 10^9
					// e.g. nano_cpus=2000000000 for 2 CPUs
					// NanoCPUs: t.GetCpus() * 1000000000,
					NanoCPUs: t.GetCpus(),
					Memory:   t.GetMemory(),
				},
				Runtime: parseRuntime(t.GetGpus()),
			},
			nil,
			"")

		if err != nil {
			atomic.AddInt64(c.nanoCPUReserved, 0-t.GetCpus())
			atomic.AddInt64(c.memoryReserved, 0-t.GetMemory())
			atomic.AddInt64(c.gpuReserved, 0-t.GetGpus())
			logger.Fatal("cannot create container instance:", err.Error())
			return t, err
		}

		t.Node = c.host
		t.ContainerId = resp.ID
		status = "created"
	}

	if status == "created" || status == "paused" {
		if err := c.ContainerStart(context.Background(), t.ContainerId, docker.ContainerStartOptions{}); err != nil {
			atomic.AddInt64(c.nanoCPUReserved, 0-t.GetCpus())
			atomic.AddInt64(c.memoryReserved, 0-t.GetMemory())
			atomic.AddInt64(c.gpuReserved, 0-t.GetGpus())
			logger.Fatal("cannot start container instance", t.GetNode()+":"+t.GetContainerId(), ":", err.Error())
			return t, err
		}

		t.StartTime = time.Now().Unix()
	}

	return t, nil
}

func (c *Client) Wait(t *types.Task) *types.Task {
	state := c.containerState(t.GetContainerId())
	if state.Status == "" {
		logger.Fatal("cannot find container", t.GetContainerId())
		return t
	}
	if state.Status == "removing" || state.Status == "exited" || state.Status == "dead" {
		t.ExitCode = int64(state.ExitCode)
		t.FinishTime = parseTime(state.FinishedAt)
		return t
	}

	if state.Status == "restarting" {
		c.Stop(t.GetContainerId())
		t.ExitCode = -9999
		t.FinishTime = parseTime(state.FinishedAt)
		return t
	}

	logger.Info("wait container", t.GetNode()+":"+t.GetContainerId(), "running")
	statusCh, errCh := c.ContainerWait(context.Background(), t.GetContainerId(), container.WaitConditionNotRunning)

	select {
	case err := <-errCh:
		if err != nil {
			logger.Fatal("container wait", t.GetNode()+":"+t.GetContainerId(), "get error signal:", err.Error())
			t.ExitCode = -9999
		}
	case status := <-statusCh:
		t.ExitCode = status.StatusCode
	}

	atomic.AddInt64(c.nanoCPUReserved, 0-t.GetCpus())
	atomic.AddInt64(c.memoryReserved, 0-t.GetMemory())
	atomic.AddInt64(c.gpuReserved, 0-t.GetGpus())

	t.FinishTime = time.Now().Unix()

	return t

}

func parseTime(t string) int64 {
	ft, err := time.Parse(time.RFC3339Nano, t)
	if err != nil {
		ft = time.Now()
	}
	return ft.Unix()
}

func parseRuntime(gpu int64) string {
	if gpu != 0 {
		return "nvidia"
	}
	return ""
}
