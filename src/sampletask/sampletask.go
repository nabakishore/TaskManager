
package sampletask

import (
	"fmt"
	"time"
	"sync"
	"../task"
	"../status"
)

// Sample Task definitions
type SampleTask task.Task

func (S *SampleTask) TRunning() bool {
	S.Lock()
        if S.Status.TaskStarted && !S.Status.TaskExited {
		S.Unlock()
                return true
        }
	S.Unlock()
        return false
}

func updateStatus(S *SampleTask) error {
	var conf status.StatusConf
	conf.WrStatus.TaskState = S.Status.TaskState
        conf.WrStatus.TaskLoopCount = S.Status.TaskLoopCount
	S.Conf.Update(conf)

	return nil
}

func getStatus(S *SampleTask) error {
	status := S.Conf.Fetch()
	S.Status.TaskState = status.TaskState
	S.Status.TaskLoopCount = status.TaskLoopCount

	return nil
}

func (S *SampleTask) TRun() {
	// Restart from a previous state if exist
	getStatus(S)

	S.Lock()
        S.Status.TaskStarted = true
	S.Unlock()
        for {
		// Update task status
		updateStatus(S)

		S.Lock()
		if S.Status.TaskExit {
			S.Unlock()
			break;
		}
		if S.Status.TaskPause {
			// Signal Pause initiator
			S.PausedCond.Broadcast()

			// Wait for Resume condition
			S.ResumeCond.Wait()

			// Task Resumed
			S.Status.TaskPaused = false
		}
		S.Status.TaskLoopCount++
		S.Unlock()

		// Sample tasklet of sleep
		time.Sleep(10 * time.Second)
        }
	S.Lock()
        S.Status.TaskExited = true
	S.ExitCond.Broadcast()
	S.Unlock()
        fmt.Println("Exited SampleTask", S.TaskName)
}

func (S *SampleTask) TPause() {
	S.Lock()

	// Set flag to pause a task
	S.Status.TaskPause = true

	// Wait for the pause
	S.PausedCond.Wait()

	S.Status.TaskPaused = true
	S.Unlock()
}

func (S *SampleTask) TResume() {
	S.Lock()
	S.Status.TaskPause = false
	S.ResumeCond.Broadcast()
	S.Unlock()
}

func (S *SampleTask) TRestart() {
	S.TExit()
	S.TRun()
}

func (S *SampleTask) TStatus() {
        fmt.Println(S.TaskName, "Running :", S.Status.TaskStarted && !S.Status.TaskExited)
	fmt.Println(S.TaskName, "Paused :", S.Status.TaskPaused)
        fmt.Println("Task loop count : ", S.Status.TaskLoopCount)
}

func (S *SampleTask) TExit() {

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
	if S.Conf != nil {
		S.Conf.Fetch()
	}
}

func NewSampleTask(name string) *SampleTask {
	var S status.StatusInterface

        t := &SampleTask {
                TaskId: name,
                TaskName: name,
        }

	t.PausedCond = sync.NewCond(t)
	t.ResumeCond = sync.NewCond(t)
	t.ExitCond = sync.NewCond(t)
	s := status.NewStatusConf(name + ".json")
	S = s
	t.Conf = S
	return t
}

