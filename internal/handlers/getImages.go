package handlers

import (
	"github.com/kumato/kumato/internal/logger"
	"io/ioutil"
	"net/http"
)

func getImages(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(dataPath + "/.images.json")
	if err != nil {
		logger.Fatal("fail to read images.json file:", err.Error())
		handleErr(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(b)
}
