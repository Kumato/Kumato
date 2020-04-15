package db

import (
	"github.com/kumato/kumato/internal/types"
)

func GetTask(t *types.Task) {
	db.First(t)
}

func GetTasks(o *types.Option) (*types.Tasks, error) {
	var ts []types.Task

	if o.GetOwnerName() == "" {
		db.Order("ID desc").
			Offset(o.GetOffset()).
			Limit(o.GetLimit()).
			Find(&ts)
	} else {
		db.Where("owner_name = ?", o.GetOwnerName()).
			Order("ID desc").
			Offset(o.GetOffset()).
			Limit(o.GetLimit()).
			Find(&ts)
	}

	re := types.Tasks{}

	for _, i := range ts {
		re.Tasks = append(re.Tasks, &types.TaskShortForm{
			Id:         i.GetId(),
			Title:      i.GetTitle(),
			ImageRepo:  i.GetImageRepo(),
			ImageTag:   i.GetImageTag(),
			CreateTime: i.GetCreateTime(),
			StartTime:  i.GetStartTime(),
			FinishTime: i.GetFinishTime(),
			ExitCode:   i.GetExitCode(),
			OwnerName:  i.GetOwnerName(),
		})
	}
	return &re, nil
}

func UpdateTask(t *types.Task) {
	db.Save(t)
}

func GetPendingTasks(r *types.Requirement) *[]types.Task {
	var t []types.Task

	db.Order("ID asc").
		Where("cpus <= ? AND gpus <= ? AND memory <= ? AND node = ? AND container_id = ?", r.GetCpus(), r.GetGpus(), r.GetMemory(), "", "").
		Find(&t)

	return &t
}

func GetRunningTasksByNode(node string) *[]types.Task {
	var t []types.Task

	db.Order("ID asc").
		Where("node = ? AND create_time > ? AND start_time > ? AND finish_time = ?", node, 0, 0, 0).
		Find(&t)

	return &t
}
