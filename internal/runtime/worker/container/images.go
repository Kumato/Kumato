package container

import (
	"context"
	docker "github.com/docker/docker/api/types"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/types"
	"sort"
	"strings"
)

func (c *Client) GetImages() (*types.Images, error) {
	imgSums, err := c.ImageList(context.Background(), docker.ImageListOptions{})

	if err != nil {
		logger.Fatal("cannot get docker image list:", err.Error())
		return &types.Images{}, err
	}

	var imgs types.Images

	imgsMap := make(map[string][]string)

	for _, r := range imgSums {
		for _, i := range r.RepoTags {
			s := strings.Split(i, ":")
			if len(s) != 2 || s[0] == "<none>" || s[1] == "<none>" {
				continue
			}
			if _, ok := imgsMap[s[0]]; !ok {
				imgsMap[s[0]] = []string{s[1]}
				continue
			}
			imgsMap[s[0]] = append(imgsMap[s[0]], s[1])
		}
	}

	for k, v := range imgsMap {
		sort.SliceStable(v, func(i, j int) bool { return v[i] < v[j] })
		imgs.ImageRepoTags = append(imgs.ImageRepoTags, &types.ImageRepoTags{
			Repo: k,
			Tags: v,
		})
	}

	return &imgs, nil
}
