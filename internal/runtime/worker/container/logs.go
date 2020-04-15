package container

import (
	"context"
	docker "github.com/docker/docker/api/types"
	"io"
)

func (c *Client) Logs(id string) (io.ReadCloser, error) {
	return c.ContainerLogs(context.Background(), id, docker.ContainerLogsOptions{ShowStdout: true})
}
