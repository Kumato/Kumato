package handlers

import (
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/types"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	o := types.Option{}

	if err := parseJSON(r.Body, &o, []string{"owner_name", "offset"}); err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}

	t, err := db.GetTasks(&o)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err)
		return
	}

	handleOK(w, t)
}
