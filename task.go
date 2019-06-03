package task

import (
	"fmt"
	"time"
)

type TaskInterface interface {
	TRunning() bool
	TRun()
	TPause()
	TRestart()
	TStatus()
	TExit()
}

type TaskResources struct {
        memory int
        cpu int
        storage int
}

type TaskStatus struct {
        TaskState int
	TaskLoopCount int
        TaskProgress int
	TaskStarted bool
	TaskExit bool
	TaskExited bool
}

type Task struct {
        TaskId string
        TaskName string
        Resource TaskResources
        Status TaskStatus
}

type TaskEngine struct {
	Name		string
	TaskCount 	int
	MaxTaskCount 	int
	Tasks		map[string] TaskInterface
}

func (S *Task) TRinning() bool {
	if S.Status.TaskStarted && !S.Status.TaskExited {
		return true
	}
	return false
} 

func (S *Task) TRun() {
	fmt.Println("Task Run")
}

func (S *Task) TPause() {
	fmt.Println("Task Pause")
}

func (S *Task) TRestart() {
	fmt.Println("Task Retart")
}

func (S *Task) TStatus() {
	fmt.Println("Task Status")
}

func (S *Task) TExit() {
	S.Status.TaskExit = true
	for ; S.Status.TaskStarted && S.Status.TaskExited != true ; {
		time.Sleep(5 * time.Second)
	}
}

func NewTask(name string) *Task {
	return &Task {
		TaskId: name,
		TaskName: name,
	}
}

func createTaskEngine(name string, maxcount int) *TaskEngine {
	return &TaskEngine{
		Name:	name,
		TaskCount: 0,
		MaxTaskCount: maxcount,
		Tasks: make(map[string] TaskInterface),
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
func (engine *TaskEngine) AddTask(name string, task TaskInterface) error {
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
	var ti TaskInterface = t
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
	var ti TaskInterface = el
	ti.TExit()
	return nil
}

// Task Status
func (engine *TaskEngine) TaskStatus(name string) error {
	el := engine.Tasks[name]
	var ti TaskInterface = el
	ti.TStatus()
	return nil
}

//List tasks
func (engine *TaskEngine) ListTask() {
	for key, _ := range engine.Tasks {
		fmt.Println(key)
	}
}
