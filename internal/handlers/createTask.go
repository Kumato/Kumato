package handlers

import (
	"errors"
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/logger"
	"github.com/kumato/kumato/internal/runtime/controller"
	"github.com/kumato/kumato/internal/types"
	"net/http"
	"regexp"
	"time"
)

func createTask(w http.ResponseWriter, r *http.Request) {
	t := types.Task{}
	oe := []string{"id",
		"description",
		"create_time",
		"start_time",
		"finish_time",
		"exit_code",
		"worker",
		"container_id",
		"gpus",
		"owner_name",
		"owner_qid",
		"owner_email",
		"node"}

	if err := parseJSON(r.Body, &t, oe); err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}

	token, err := getToken(r)
	if err != nil {
		handleErr(w, http.StatusForbidden, err)
		return
	}

	re := regexp.MustCompile(`^[0-9a-zA-Z .]+$`)

	if !(re.MatchString(t.GetTitle()) && (re.MatchString(t.GetDescription()) || t.GetDescription() == "")) {
		handleErr(w, http.StatusBadRequest, errors.New("title or description contains invalid characters"))
		return
	}

	t.CreateTime = time.Now().Unix()
	t.OwnerName = token.Name
	t.OwnerEmail = token.Email
	t.OwnerQid = token.Qid

	logger.Warn("create task:", t)
	db.UpdateTask(&t)

	go controller.AssignTask()

	handleOK(w, t)
}
