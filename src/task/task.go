package task

import (
	"fmt"
	"sync"
	"../status"
)

type TaskInterface interface {
	TRunning() bool
	TRun()
	TPause()
	TResume()
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
	TaskPause bool
	TaskPaused bool
}

type Task struct {
	sync.Mutex
        TaskId string
        TaskName string
        Resource TaskResources
        Status TaskStatus
	PausedCond *sync.Cond
	ResumeCond *sync.Cond
	ExitCond *sync.Cond
	Conf status.StatusInterface
}

func (S *Task) TRunning() bool {
	S.Lock()
	if S.Status.TaskStarted && !S.Status.TaskExited {
		S.Unlock()
		return true
	}
	S.Unlock()
	return false
} 

func updateStatus(S *Task) error {
	var conf status.StatusConf
	conf.WrStatus.TaskState = S.Status.TaskState
	conf.WrStatus.TaskLoopCount = S.Status.TaskLoopCount
	S.Conf.Update(conf)
	return nil
}

func (S *Task) TRun() {
	fmt.Println("Task Run")
	S.Lock()
	S.Status.TaskStarted = true
	S.Unlock()

	for {
		S.Lock()
		if S.Status.TaskExit {
			S.Unlock()
			break
		}
		if S.Status.TaskPause {
			S.PausedCond.Broadcast()
			S.ResumeCond.Wait()
			S.Status.TaskPaused = false
		}
		S.Status.TaskLoopCount++
		S.Unlock()

		// Do something

	}

	S.Lock()
	S.Status.TaskExited = true
	S.ExitCond.Broadcast()
	S.Unlock()
}

func (S *Task) TPause() {
	S.Lock()
	S.Status.TaskPause = true
	S.PausedCond.Wait()
	S.Status.TaskPaused = true
	S.Unlock()
	fmt.Println("Task Pause")
}

func (S *Task) TResume() {
	S.Lock()
	S.Status.TaskPause = false
	S.ResumeCond.Broadcast()
	S.Unlock()
	fmt.Println("Task Resume")
}

func (S *Task) TRestart() {
	fmt.Println("Task Retart")
	S.TExit()
	S.TRun()
}

func (S *Task) TStatus() {
	fmt.Println("Task Status")
}

func (S *Task) TExit() {
	S.Lock()
	S.Status.TaskExit = true
	S.ResumeCond.Broadcast()
	S.ExitCond.Wait()
	S.Status.TaskStarted = false
	S.Status.TaskExit = false
	S.Status.TaskExited = false
	S.Status.TaskPause = false
	S.Status.TaskPaused = false
	S.Unlock()

}

func NewTask(name string) *Task {
	t := &Task {
		TaskId: name,
		TaskName: name,
	}
	t.PausedCond = sync.NewCond(t)
	t.ResumeCond = sync.NewCond(t)
	t.ExitCond = sync.NewCond(t)
	conf := status.NewStatusConf(name + ".json")
	t.Conf = conf
	return t
}
