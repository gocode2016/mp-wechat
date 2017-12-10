package handler

import (
	"net/http"

	"github.com/golang/glog"
)

func root(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	glog.Info("==> receive from: %v\n", r.RemoteAddr)
	glog.Info("==> req header:%v", r.Header)

	w.Write([]byte("hello 微信"))
}
