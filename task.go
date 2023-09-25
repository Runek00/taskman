package main

import "fmt"

var nextTaskId int = 0
var taskStore map[int]Task = make(map[int]Task)

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
	taskStore[task.id] = task
	return task
}

func showTasks() {
	for _, t := range taskStore {
		fmt.Printf("\nID: %v	Title: %v	Completed: %v", t.id, t.title, t.completed)
	}
}

func completeTask(id int) {
	t := taskStore[id]
	t.completed = true
	taskStore[id] = t
}

func deleteTask(id int) {
	delete(taskStore, id)
}

func main() {
	addTask("q")
	addTask("w")
	addTask("e")
	addTask("r")
	addTask("t")
	addTask("y")
	completeTask((3))
	deleteTask(4)
	showTasks()
}
