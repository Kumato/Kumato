package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func RegisterHandlers(basePath string, mux *http.ServeMux, data string) {
	basePath = filepath.Clean(basePath)
	if data == "" {
		data = os.TempDir()
	}

	dataPath = data + "/"
	uploadPath = dataPath + "file/"

	mux.Handle(basePath+"/getTask", secure(getTask))
	mux.Handle(basePath+"/createTask", secure(createTask))
	mux.Handle(basePath+"/stopTask", secure(stopTask))
	mux.Handle(basePath+"/getTasks", secure(getTasks))

	mux.Handle(basePath+"/getImages", secure(getImages))
	mux.Handle(basePath+"/uploadFile", secure(uploadFile))
	mux.Handle(basePath+"/getFile/", secure(getFile))
	mux.Handle(basePath+"/internal/getFile/", secureInternal(getFile))
	mux.Handle(basePath+"/internal/uploadLog", secureInternal(uploadLog))
	mux.Handle(basePath+"/internal/uploadResult", secureInternal(uploadResult))

	mux.HandleFunc(basePath+"/login", login)
	mux.Handle(basePath+"/getMe", secure(getMe))
	mux.Handle(basePath+"/getUsers", secure(getUsers))
	mux.Handle(basePath+"/getSysLoad", secure(getSysLoad))
	mux.Handle(basePath+"/getAlert", secure(getAlert))
}
