package handlers

import (
	"errors"
	"github.com/kumato/kumato/internal/logger"
	"io/ioutil"
	"net/http"
	"path"
)

func uploadResult(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(32 << 20)
	//f, handler, err := r.FormFile("file")
	//if err != nil {
	//	handleErr(w, http.StatusBadRequest, err)
	//	return
	//}
	//defer f.Close()
	//
	//nf, err := os.OpenFile(uploadPath+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0644)
	//if err != nil {
	//	handleErr(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//defer nf.Close()
	//
	//if _, err := io.Copy(nf, f); err != nil {
	//	handleErr(w, http.StatusInternalServerError, err)
	//	return
	//}
	//
	//logger.Info("file created:", handler.Filename)

	id := r.Header.Get("RESULT_FILE")
	if id == "" || path.Ext(id) != ".zip" {
		handleErr(w, http.StatusBadRequest, errors.New("RESULT_FILE is not set"))
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
}
