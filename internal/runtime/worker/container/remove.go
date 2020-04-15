package container

import (
	"context"
	docker "github.com/docker/docker/api/types"
	"github.com/kumato/kumato/internal/logger"
)

func (c *Client) Remove(id string) {
	if err := c.ContainerRemove(context.Background(), id, docker.ContainerRemoveOptions{}); err != nil {
		logger.Fatal("cannot remove container", id, ":", err.Error())
	}
}
