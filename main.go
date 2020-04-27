package main

import (
	"github.com/Jopoleon/geekbrains_0/app"
	"github.com/Jopoleon/geekbrains_0/config"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	ll := logrus.New()
	childFormatter := logrus.JSONFormatter{}
	runtimeFormatter := &runtime.Formatter{ChildFormatter: &childFormatter}
	runtimeFormatter.Line = true
	runtimeFormatter.File = true

	ll.SetFormatter(runtimeFormatter)

	cfg := config.NewConfig()

	a, err := app.New(cfg, ll)
	if err != nil {
		logrus.Fatal(err)
	}
	a.Run()
}
