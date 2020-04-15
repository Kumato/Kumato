package handlers

import (
	"crypto/md5"
	"fmt"
	"github.com/kumato/kumato/internal/logger"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	f, handler, err := r.FormFile("file")
	if err != nil {
		handleErr(w, http.StatusBadRequest, err)
		return
	}
	defer f.Close()

	normalizedName := randomName()
	for {
		if _, err := os.Stat(uploadPath + "/" + normalizedName); os.IsNotExist(err) {
			break
		}
		normalizedName = randomName()
	}

	nf, err := os.OpenFile(uploadPath+"/"+normalizedName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		handleErr(w, http.StatusInternalServerError, err)
		return
	}

	defer nf.Close()

	h := md5.New()

	writers := io.MultiWriter(h, nf)

	if _, err := io.Copy(writers, f); err != nil {
		handleErr(w, http.StatusInternalServerError, err)
		return
	}

	id := fmt.Sprintf("%x%s", h.Sum(nil), path.Ext(handler.Filename))

	if err := os.Rename(uploadPath+"/"+normalizedName, uploadPath+"/"+id); err != nil {
		handleErr(w, http.StatusInternalServerError, err)
		return
	}

	logger.Info("file created:", id)

	handleOK(w, struct {
		ID string `json:"id"`
	}{id})
}

func randomName() string {
	b := make([]byte, 64)
	rand.Seed(time.Now().UnixNano())
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
