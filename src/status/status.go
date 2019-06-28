package status

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

type StatustoWrite struct {
        TaskState int
        TaskLoopCount int
        TaskProgress int
}
type StatusConf struct {
        Handle *os.File
        Name string
        WrStatus StatustoWrite
}

type StatusInterface interface {
	Configure() error 
	Update(StatusConf) error
	Fetch() StatustoWrite 
	Close() error
}

func (conf *StatusConf) Configure() error {
	_, err := conf.Handle.Stat()

	return err
}

func (conf *StatusConf) Update(S StatusConf) error {
	conf.WrStatus.TaskState = S.WrStatus.TaskState
	conf.WrStatus.TaskLoopCount = S.WrStatus.TaskLoopCount
	file, _ := json.MarshalIndent(conf.WrStatus, "", " ")
	err := ioutil.WriteFile(conf.Name, file, 0644) 
	return err
}

func (conf *StatusConf) Fetch() StatustoWrite {
	data := StatustoWrite {}
	file, err := ioutil.ReadFile(conf.Name)
	if err != nil {
		return data
	}
	_= json.Unmarshal([]byte(file), &data)
	fmt.Println(data)
	return data 
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
		WrStatus: StatustoWrite {
				0,
				0,
				0,
			},
	}
	s.Handle,_ = os.Open(name)

	return s
}

