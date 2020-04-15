package handlers

import (
	"io/ioutil"
	"net/http"
)

func getAlert(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(dataPath + "/.alert.json")
	if err != nil {
		handleOK(w, nil)
		return
	}
	w.Write(b)
}