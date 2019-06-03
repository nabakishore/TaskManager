
package sampletask

import (
	"fmt"
	"time"
	"../task"
)

// Sample Task definitions
type SampleTask task.Task

func (S *SampleTask) TRunning() bool {
        if S.Status.TaskStarted && !S.Status.TaskExited {
                return true
        }

        return false
}

func (S *SampleTask) TRun() {
        S.Status.TaskStarted = true
        for ; S.Status.TaskExit != true ; {
                S.Status.TaskLoopCount++
                time.Sleep(10 * time.Second)
        }
        S.Status.TaskExited = true
        fmt.Println("Exited SampleTask", S.TaskName)
}

func (S *SampleTask) TPause() {

}

func (S *SampleTask) TRestart() {

}

func (S *SampleTask) TStatus() {
        fmt.Println(S.TaskName, "Running :", S.Status.TaskStarted && !S.Status.TaskExited)
        fmt.Println("Task loop count : ", S.Status.TaskLoopCount)
}

func (S *SampleTask) TExit() {
        S.Status.TaskExit = true
        for ; S.Status.TaskStarted && S.Status.TaskExited != true ; {
                time.Sleep(5 * time.Second)
        }
        S.Status.TaskStarted = false
        S.Status.TaskExit = false
        S.Status.TaskExited = false
}

func NewSampleTask(name string) *SampleTask {
        return &SampleTask {
                TaskId: name,
                TaskName: name,
        }
}

