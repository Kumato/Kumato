package handlers

import (
	"github.com/kumato/kumato/internal/db"
	"github.com/kumato/kumato/internal/types"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	o := types.Option{}

	if err := parseJSON(r.Body, &o, []string{"offset", "owner_name"}); err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}

	users := db.GetUsers(&o)

	for i := 0; i < len(users); i++ {
		users[i].Qid = ""
	}

	handleOK(w, users)
}
