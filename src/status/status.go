package status

import (
	"os"
)

type StatusInterface interface {
	Configure() error 
	Update() error
	Fetch() error
	Close() error
}

type StatusConf struct {
	Handle *os.File
	Name string
	
}

func (conf *StatusConf) Configure() error {
	_, err := conf.Handle.Stat()

	return err
}

func (conf *StatusConf) Update() error {
	_, err := conf.Handle.Stat()

	return err
}

func (conf *StatusConf) Fetch() error {
	_, err := conf.Handle.Stat()

	return err
}

func (conf *StatusConf) Close() error {
	if conf.Handle != nil {
		conf.Handle.Close()
	}

	return nil
}

func NewStatusConf(name string) *StatusConf {
	s := &StatusConf {
		Name: "/tmp/" + name,
	}
	s.Handle,_ = os.Open(name)

	return s
}

