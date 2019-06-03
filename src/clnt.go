package main

import (
	"fmt"
	"time"
	"./task"
)


func Usage() {
	fmt.Println("\t[help] : help")
	fmt.Println("\t[start] : Start Task Engine")
	fmt.Println("\t[createtask <name>] : Create a new task")
	fmt.Println("\t[deletetask <name>] : Delete task")
	fmt.Println("\t[runtask <name>] : Run a given task")
	fmt.Println("\t[list] : List tasks")
	fmt.Println("\t[status <name>] : Get status")
	fmt.Println("\t[stop <name>] : Stop a task")
	fmt.Println("\t[exit] : exit Task Engine")
}

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

func main() {
	var s string
	var e *task.TaskEngine


	for ; s != "exit" ; {
		fmt.Print(">")

		fmt.Scanf("%s", &s)

		switch s {
		case "help":
			Usage()
		case "start":
			fmt.Println("Start task Manager")
			e,_ = task.NewTaskEngine("Test Engine", 200)
		case "createtask":
			var taskname string
			fmt.Scanf("%s", &taskname)
			sampletask := NewSampleTask(taskname)
			var S task.TaskInterface
			S = sampletask
			e.AddTask(taskname, S)
		case "deletetask":
			var taskname string
			fmt.Scanf("%s", &taskname)
			e.RemoveTask(taskname)
		case "runtask":
			var name string
			fmt.Scanf("%s", &name)
			e.RunTask(name)
		case "list":
			e.ListTask()
		case "status":
			var name string
			fmt.Scanf("%s", &name)
			e.TaskStatus(name)
		case "stop":
			var name string
			fmt.Scanf("%s", &name)
			e.StopTask(name)
		case "exit":
			e.Exit()
			fmt.Println("exit")
		default:
			fmt.Println("Invalid command")
			Usage()
		}
	}
	fmt.Println("Exited")
}
