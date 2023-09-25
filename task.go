package main

var idSeq int = 0
var taskStore []Task = make([]Task, 0)

type Task struct {
	id        int
	title     string
	completed bool
}

func createTask(title string) Task {
	out := Task{idSeq, title, false}
	idSeq++
	return out
}

func addTask(title string) {
	taskStore = append(taskStore, createTask(title))
}
