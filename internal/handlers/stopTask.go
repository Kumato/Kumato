package handlers

import (
	"errors"
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/runtime/controller"
	"github.com/kumato/kumato/internal/types"
	"net/http"
	"time"
)

func stopTask(w http.ResponseWriter, r *http.Request) {
	i := IDOnlyRequest{}

	if err := parseJSON(r.Body, &i, []string{}); err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}

	t := types.Task{Id: i.ID}

	db.GetTask(&t)

	token, err := getToken(r)
	if err != nil {
		handleErr(w, http.StatusForbidden, err)
		return
	}

	if t.GetOwnerName() != token.Name || t.GetOwnerEmail() != token.Email || t.GetOwnerQid() != token.Qid {
		handleErr(w, http.StatusForbidden, errors.New("attempt to stop a task which belongs to other user"))
		return
	}

	err = controller.StopTask(t.Node, t.ContainerId)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err)
		return
	}

	time.Sleep(2 * time.Second)

	handleOK(w, responseMessage{"ok", ""})
}
