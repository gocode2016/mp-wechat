package config

import (
	"fmt"

	"github.com/golang/glog"

	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	Token string `yaml:"token"`
}

func (c Configuration) String() string {
	return fmt.Sprintf("token=%s", c.Token)
}

var Config = &Configuration{}

func init() {
	b, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		glog.Fatalf("read config file error:%v", err)
	}
	err = yaml.Unmarshal(b, Config)

	if err != nil {
		glog.Fatalf("parse config file error,err=%v,file=%v", err, string(b))
	}
}

func DumpConfig() {
	glog.Infof("config:\n%v", Config)
}
