package taskengine

import (
	"fmt"
	"errors"
	"../task"
	"../status"
)

type TaskEngine struct {
        Name            string
        TaskCount       int
        MaxTaskCount    int
        Tasks           map[string] task.TaskInterface
	StatusConfig	map[string] status.StatusInterface
}

func configureStatus(statusdb string) *status.StatusConf {

	conf := new(status.StatusConf)
	conf.Name = statusdb
	conf.Configure()
	return conf
}

func createTaskEngine(name string, maxcount int) *TaskEngine {
        return &TaskEngine{
                Name:   name,
                TaskCount: 0,
                MaxTaskCount: maxcount,
                Tasks: make(map[string] task.TaskInterface),
		StatusConfig: make(map[string] status.StatusInterface),
        }
}

// Start Task Engine
func NewTaskEngine(name string, maxtask int) (*TaskEngine, error) {
        eng := createTaskEngine(name, maxtask)

        return eng, nil
}

// Exit Tasks in thas engine
func (engine *TaskEngine) Exit() {
        for _, el := range engine.Tasks {
		if el.TRunning() {
	                el.TExit()
		}
        }
}

// Add task to the task map
func (engine *TaskEngine) AddTask(name string, task task.TaskInterface) error {
	var S status.StatusInterface
        engine.Tasks[name] = task
	s := status.NewStatusConf(name)
	S = s
	engine.StatusConfig[name] = S 
	S.Configure()
        return nil
}

// Remove task
func (engine *TaskEngine) RemoveTask(name string) error {
	s, ok := engine.StatusConfig[name]
	if !ok {
		return errors.New("No entry")
	}
	var si status.StatusInterface = s
	si.Close()
	delete(engine.StatusConfig, name)
        delete(engine.Tasks, name)

	return nil

}


// Run a task
func (engine *TaskEngine) RunTask(name string) error {
        t, ok := engine.Tasks[name]
	if !ok {
		return errors.New("No entry")
	}
        var ti task.TaskInterface = t
        if ti.TRunning() {
                fmt.Println("Task already running")
                return nil
        }
        go ti.TRun()
        return nil
}

// Stop a task
func (engine *TaskEngine) StopTask(name string) error {
        el, ok := engine.Tasks[name]
	if !ok {
		return errors.New("No entry")
	}
        var ti task.TaskInterface = el
        ti.TExit()
        return nil
}

// Task Status
func (engine *TaskEngine) TaskStatus(name string) error {
        el, ok := engine.Tasks[name]
	if !ok {
		return errors.New("No entry") 
	}
	st, ok := engine.StatusConfig[name]
	if !ok {
		return errors.New("No entry")
	}
        var ti task.TaskInterface = el
	var si status.StatusInterface = st
	fmt.Println(si)
        ti.TStatus()

        return nil
}

//List tasks
func (engine *TaskEngine) ListTask() {
        for key, _ := range engine.Tasks {
                fmt.Println(key)
        }
}

// Pause task
func (engine *TaskEngine) PauseTask(name string) error {
	el, ok := engine.Tasks[name]
	if !ok {
		return errors.New("No entry")
	}
	var ti task.TaskInterface = el
	ti.TPause()
	return nil
}

// Resume task
func (engine *TaskEngine) ResumeTask(name string) error {
	el, ok := engine.Tasks[name]
	if !ok {
		return errors.New("No entry")
	}
	var ti task.TaskInterface = el
	ti.TResume()
	return nil
}
