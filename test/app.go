package main

import (
	"github.com/dynamicgo/go-config"
	"github.com/dynamicgo/slf4go"
	"github.com/gomeshnetwork/gomesh"
	"github.com/gomeshnetwork/gomesh/app"
	"github.com/gomeshnetwork/localservice"
)

var logger = slf4go.Get("test")

type serviceA struct {
	B    *serviceB `inject:"B"`
	Name string
}

func (a *serviceA) Start() error {
	logger.InfoF("A:%s, B:%s", a.Name, a.B.Name)
	return nil
}

type serviceB struct {
	A    *serviceA `inject:"A"`
	Name string
}

func (b *serviceB) Start() error {
	logger.InfoF("B:%s, A:%s", b.Name, b.A.Name)
	return nil
}

func main() {

	localservice.Register("A", func(config config.Config) (gomesh.Service, error) {
		return &serviceA{
			Name: config.Get("Name").String(""),
		}, nil
	})

	localservice.Register("B", func(config config.Config) (gomesh.Service, error) {
		return &serviceB{
			Name: config.Get("Name").String(""),
		}, nil
	})

	app.Run("test")
}
