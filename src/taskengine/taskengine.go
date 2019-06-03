package taskengine

import (
	"fmt"
	"../task"
)

type TaskEngine struct {
        Name            string
        TaskCount       int
        MaxTaskCount    int
        Tasks           map[string] task.TaskInterface
}

func createTaskEngine(name string, maxcount int) *TaskEngine {
        return &TaskEngine{
                Name:   name,
                TaskCount: 0,
                MaxTaskCount: maxcount,
                Tasks: make(map[string] task.TaskInterface),
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
                el.TExit()
        }
}

// Add task to the task map
func (engine *TaskEngine) AddTask(name string, task task.TaskInterface) error {
        engine.Tasks[name] = task

        return nil
}

// Remove task
func (engine *TaskEngine) RemoveTask(name string) {
        delete(engine.Tasks, name)
}


// Run a task
func (engine *TaskEngine) RunTask(name string) error {
        t := engine.Tasks[name]
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
        el := engine.Tasks[name]
        var ti task.TaskInterface = el
        ti.TExit()
        return nil
}

// Task Status
func (engine *TaskEngine) TaskStatus(name string) error {
        el := engine.Tasks[name]
        var ti task.TaskInterface = el
        ti.TStatus()
        return nil
}

//List tasks
func (engine *TaskEngine) ListTask() {
        for key, _ := range engine.Tasks {
                fmt.Println(key)
        }
}

