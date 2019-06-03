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
