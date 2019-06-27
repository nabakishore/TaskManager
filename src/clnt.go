package main

import (
	"fmt"
	"./task"
	"./sampletask"
	"./taskengine"
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
	fmt.Println("\t[pause <name>] : Pause a task")
	fmt.Println("\t[resume <name>] : Resume a paused task")
	fmt.Println("\t[exit] : exit Task Engine")
}

func main() {
	var s string
	var e *taskengine.TaskEngine

	for ; s != "exit" ; {
		fmt.Print(">")

		fmt.Scanf("%s", &s)

		switch s {
		case "help":
			Usage()
		case "start":
			fmt.Println("Start task Manager")
			e,_ = taskengine.NewTaskEngine("Test Engine", 200)
		case "createtask":
			var taskname string
			var S task.TaskInterface
			fmt.Scanf("%s", &taskname)
			t := sampletask.NewSampleTask(taskname)
			S = t
			if e != nil {
				e.AddTask(taskname, S)
			}
		case "deletetask":
			var taskname string
			fmt.Scanf("%s", &taskname)
			if e != nil {
				e.RemoveTask(taskname)
			}
		case "runtask":
			var name string
			fmt.Scanf("%s", &name)
			if e != nil {
				e.RunTask(name)
			}
		case "list":
			if e != nil {
				e.ListTask()
			}
		case "status":
			var name string
			fmt.Scanf("%s", &name)
			if e != nil {
				err := e.TaskStatus(name)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "stop":
			var name string
			fmt.Scanf("%s", &name)
			if e != nil {
				e.StopTask(name)
			}
		case "pause":
			var name string
			fmt.Scanf("%s", &name)
			if e != nil {
				e.PauseTask(name)
			}
		case "resume":
			var name string
			fmt.Scanf("%s", &name)
			if e != nil {
				e.ResumeTask(name)
			}
		case "exit":
			if e != nil {
				e.Exit()
			}
			fmt.Println("exit")
		default:
			fmt.Println("Invalid command")
			Usage()
		}
	}
	fmt.Println("Exited")
}
