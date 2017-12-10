package main

import (
	"flag"
	"net/http"
	"rw/mp-wechat/config"
	"rw/mp-wechat/handler"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	config.DumpConfig()
	handler.Register()
	glog.Info("==>server start:")
	glog.Fatal(http.ListenAndServe(":80", nil))
}
