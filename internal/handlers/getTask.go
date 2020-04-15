package handlers

import (
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/types"
	"net/http"
)

func getTask(w http.ResponseWriter, r *http.Request) {
	i := IDOnlyRequest{}

	if err := parseJSON(r.Body, &i, []string{}); err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}

	t := types.Task{Id: i.ID}

	db.GetTask(&t)

	logger.Info("get task:", t)

	t.OwnerQid = ""
	handleOK(w, t)
}
