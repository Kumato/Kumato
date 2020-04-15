package controller

import (
	"context"
	"errors"
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/types"
	"reflect"
	"sync"
)

var (
	nodes      sync.Map
	assignWork *sync.Once
)

func init() {
	assignWork = new(sync.Once)
}

func StopTask(node, containerID string) error {
	v, ok := nodes.Load(node)
	if ok {
		_, e := v.(types.WorkerClient).StopTask(context.Background(), &types.Task{ContainerId: containerID})
		return e
	}

	return errors.New("node " + node + "has not yet registered")
}

func AssignTask() {
	assignWork.Do(func() {
		logger.Warn("start task assignment ...")
		nodes.Range(nodesRange)
		logger.Warn("task assignment finished")
		assignWork = new(sync.Once)
		logger.Warn("task assignment lock reset")
	})
}

func nodesRange(key, value interface{}) bool {
	for _, i := range *db.GetRunningTasksByNode(key.(string)) {
		t, err := value.(types.WorkerClient).RunTask(context.Background(), &i)
		if err != nil {
			logger.Fatal("fail to assign task to", key, ":", err.Error())
			return true
		}
		if !reflect.DeepEqual(*t, i) {
			logger.Info("found task", t.GetId(), "different between controller and worker")
			db.UpdateTask(t)
		}
	}

	stat, err := value.(types.WorkerClient).GetStats(context.Background(), &types.Empty{})
	if err != nil {
		logger.Fatal("fail to get stats from", key, ":", err.Error())
		return true
	}

	r := types.Requirement{
		Cpus:   stat.NanoCpuTotal - stat.NanoCpuUsed,
		Gpus:   stat.GpuTotal - stat.GpuUsed,
		Memory: stat.MemoryTotal - stat.MemoryUsed,
	}

	for _, i := range *db.GetPendingTasks(&r) {
		if i.GetId() == 0 {
			return true
		}

		t, err := value.(types.WorkerClient).RunTask(context.Background(), &i)
		if err != nil {
			logger.Fatal("fail to assign task to", key, ":", err.Error())
			return true
		}

		db.UpdateTask(t)
		logger.Warn("node", t.GetNode(), "obtained task:", t)

		r.Cpus -= t.Cpus
		r.Gpus -= t.Gpus
		r.Memory -= t.Memory
	}

	return true
}

func SysLoad() [][]types.Stats {
	var s, down [][]types.Stats

	nodes.Range(func(k, v interface{}) bool {
		load, err := v.(types.WorkerClient).GetLoad(context.Background(), &types.Empty{})
		if err == nil {
			stat, err := v.(types.WorkerClient).GetStats(context.Background(), &types.Empty{})
			if err == nil {
				i := make([]types.Stats, 2)
				i[0] = *load
				i[1] = *stat
				s = append(s, i)
				return true
			}
		}
		down = append(down, []types.Stats{{Hostname: k.(string)}})
		return true
	})

	return append(s, down...)
}

//func GetImages() []types.ImageRepoTags {
//	var (
//		res []types.ImageRepoTags
//		count int
//	)
//
//	nodes.Range(func(k, v interface{}) bool {
//		imgs, err := v.(types.WorkerClient).GetImages(context.Background(), &types.Empty{})
//		if err == nil {
//			for _, i := range imgs.GetImageRepoTags() {
//				res = append(res, *i)
//			}
//			count++
//		}
//		return true
//	})
//
//
//}
