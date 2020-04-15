package handlers

import (
	"errors"
	"github.com/kumato/kumato/internal/logger"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func uploadLog(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("LOG_FILE")
	if id == "" || path.Ext(id) != ".log" {
		handleErr(w, http.StatusBadRequest, errors.New("LOG_FILE is not set"))
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}

	if err := ioutil.WriteFile(uploadPath+"/"+id, b, 0644); err != nil {
		handleErr(w, http.StatusInternalServerError, err)
		return
	}

	logger.Info("file created:", id)

	os.Remove(uploadPath + "/" + r.Header.Get("FILE_URI"))
}
