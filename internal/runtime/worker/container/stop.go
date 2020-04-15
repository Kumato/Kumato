package container

import (
	"context"
	"time"
)

func (c *Client) Stop(id string) error {
	var duration time.Duration
	return c.ContainerStop(context.Background(), id, &duration)
}
