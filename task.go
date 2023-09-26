package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var nextTaskId int = 0
var taskStore map[int]Task = make(map[int]Task)

type Task struct {
	Id        int
	Title     string
	Completed bool
}

func createTask(title string) Task {
	out := Task{nextTaskId, title, false}
	nextTaskId++
	return out
}

func addTask(title string) Task {
	task := createTask(title)
	taskStore[task.Id] = task
	return task
}

func showTasks() {
	for _, t := range taskStore {
		fmt.Printf("\nID: %v	Title: %v	Completed: %v", t.Id, t.Title, t.Completed)
	}
}

func completeTask(id int) {
	t, ok := taskStore[id]
	if !ok {
		fmt.Print("No such task")
		return
	}
	t.Completed = true
	taskStore[id] = t
}

func deleteTask(id int) {
	_, ok := taskStore[id]
	if !ok {
		fmt.Print("No such task")
		return
	}
	delete(taskStore, id)
}

func interactive() {
	load()
	for {
		fmt.Print("\nWhat do you want to do?\n(add) add task\n(del) delete task\n(com) complete task\n(show) show tasks\n(exit) exit\n")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		switch strings.TrimSpace(input) {
		case "add":
			addTask(getTitle())
		case "del":
			deleteTask(getId())
		case "com":
			completeTask(getId())
		case "show":
			showTasks()
		case "exit":
			fmt.Print("BAIII!\n")
			save()
			return
		default:
			fmt.Print("No such option\n")
		}
	}
}

func getTitle() string {
	fmt.Print("\nWhat is the title?\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(input)
}

func getId() int {
	fmt.Print("\nWhat is the id?\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return -1
	}
	inn, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return inn
}

func save() {
	marshaled, err := json.Marshal(taskStore)
	if err != nil {
		fmt.Print(err)
		return
	}
	err = os.WriteFile("store.json", marshaled, 0777)
	if err != nil {
		fmt.Print(err)
		return
	}
}

func load() {
	marshaled, err := os.ReadFile("store.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	err = json.Unmarshal(marshaled, &taskStore)
	if err != nil {
		fmt.Print(err)
		return
	}
	for k := range taskStore {
		if k >= nextTaskId {
			nextTaskId = k + 1
		}
	}

}

func main() {
	// For testing
	// addTask("q")
	// addTask("w")
	// addTask("e")
	// addTask("r")
	// addTask("t")
	// addTask("y")
	// completeTask((3))
	showTasks()
	interactive()
}
