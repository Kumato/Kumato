package handlers

import (
	"github.com/kumato/kumato/internal/runtime/controller"
	"net/http"
)

func getSysLoad(w http.ResponseWriter, r *http.Request) {
	handleOK(w, controller.SysLoad())
}
