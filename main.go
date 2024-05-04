package main

import (
	"github.com/seyio91/kube-deprecated-apis/pkg/app"
)

func main() {
	app, err := app.New()

	if err != nil {
		panic(err)
	}

	if err = app.Run(); err != nil {
		panic(err)
	}
}
