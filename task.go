package main

import "fmt"

var nextTaskId int = 0
var taskStore []Task = make([]Task, 0)

type Task struct {
	id        int
	title     string
	completed bool
}

func createTask(title string) Task {
	out := Task{nextTaskId, title, false}
	nextTaskId++
	return out
}

func addTask(title string) Task {
	task := createTask(title)
	taskStore = append(taskStore, task)
	return task
}

func showTasks() {
	for _, t := range taskStore {
		fmt.Printf("\nID: %v	Title: %v	Completed: %v", t.id, t.title, t.completed)
	}
}

func completeTask(id int) {
	for i, t := range taskStore {
		if t.id == id {
			t.completed = true
			taskStore[i] = t
			break
		}
	}
}

func main() {
	addTask("q")
	addTask("w")
	addTask("e")
	addTask("r")
	addTask("t")
	addTask("y")
	completeTask((3))
	showTasks()
}
