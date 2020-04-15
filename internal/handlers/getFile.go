package handlers

import (
	"io/ioutil"
	"net/http"
	"path"
)

func getFile(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(uploadPath + "/" + path.Base(r.URL.Path))
	if err != nil {
		handleErr(w, http.StatusNotFound, err)
	}

	w.Write(b)
}
